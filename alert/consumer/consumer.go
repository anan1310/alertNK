package consumer

import (
	"alarm_collector/alert/process"
	"alarm_collector/alert/sender"
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/internal/models/system"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/hash"
	"alarm_collector/pkg/utils/templates"
	"time"

	"sync"
)

//TODO:告警->消费者

type Consume struct {
	ctx *ctx.Context
	sync.RWMutex
	models.AlertCurEvent
	// 从 Redis 中读取当前告警事件提到内存做处理.
	alertsMap map[string][]models.AlertCurEvent
	// 告警分组
	preStoreAlertGroup map[string][]models.AlertCurEvent
	Timing             map[string]int
}

type InterEvalConsume interface {
	Run()
}

func NewInterEvalConsumeWork(ctx *ctx.Context) InterEvalConsume {
	return &Consume{
		ctx:                ctx,
		alertsMap:          make(map[string][]models.AlertCurEvent),
		preStoreAlertGroup: make(map[string][]models.AlertCurEvent),
		Timing:             make(map[string]int),
	}
}

// Run 启动告警消费进程
func (ec *Consume) Run() {

	action := func() {
		//获取缓存所有Firing的Keys
		alertsCurEventKeys := process.GetRedisFiringKeys(ec.ctx)
		for _, key := range alertsCurEventKeys {
			//告警源信息
			alert := ec.ctx.Redis.Event().GetCache(key)
			// 过滤空指纹告警
			if alert.Fingerprint == "" {
				continue
			}
			//将告警源提取到alertsMap
			ec.addAlertToRuleIdMap(alert)
		}

		for key, alerts := range ec.alertsMap {
			if len(alerts) == 0 {
				continue
			}

			// 计算告警组的等待时间
			var waitTime int
			alert := ec.ctx.Redis.Event().GetCache(key)
			if alert.LastSendTime == 0 {
				// 如果是初次告警, 那么等当前告警组时间到达 groupWait 的时间则推送告警
				waitTime = global.Config.Server.GroupWait
			} else {
				// 当前告警组时间到达 groupInterval 的时间则推送告警
				waitTime = global.Config.Server.GroupInterval
			}
			if ec.Timing[key] >= waitTime {
				curEvent := ec.filterAlerts(ec.alertsMap[key])
				ec.fireAlertEvent(curEvent)
				// 执行一波后 必须重新清空alerts组中的数据。
				ec.clear(key)
			}
			ec.Timing[key]++
		}
	}
	//告警推送 每隔多长时间推送一次 默认是一秒
	ticker := time.Tick(time.Second)

	go func() {
		for range ticker {
			action()
		}
	}()

}

// 过滤告警
func (ec *Consume) filterAlerts(alerts []models.AlertCurEvent) map[string][]models.AlertCurEvent {

	var newAlertsMap = make(map[string][]models.AlertCurEvent)

	// 根据相同指纹进行去重
	newAlert := ec.removeDuplicates(alerts)
	// 将通过指纹去重后以Fingerprint为Key的Map转换成以原来RuleName为Key的Map (同一告警类型聚合)
	for _, alert := range newAlert {
		// 重复通知，如果是初次推送不用进一步判断。
		if !alert.IsRecovered {
			if alert.LastSendTime == 0 || alert.LastEvalTime >= alert.LastSendTime+alert.RepeatNoticeInterval*60 {
				newAlertsMap[alert.RuleName] = append(newAlertsMap[alert.RuleName], alert)
			}
		}
		if alert.IsRecovered {
			newAlertsMap[alert.RuleName] = append(newAlertsMap[alert.RuleName], alert)
		}
	}

	return newAlertsMap

}

// 告警事件提取到内存中
func (ec *Consume) addAlertToRuleIdMap(alert models.AlertCurEvent) {
	// 锁定
	ec.Lock()
	defer ec.Unlock()
	// 直接替换或添加告警
	ec.alertsMap[alert.RuleId] = []models.AlertCurEvent{alert}
}

// 清除本地缓存
func (ec *Consume) clear(ruleId string) {

	for key := range ec.alertsMap {
		delete(ec.alertsMap, key)
	}
	for key := range ec.preStoreAlertGroup {
		delete(ec.preStoreAlertGroup, key)
	}
	ec.Timing[ruleId] = 0

}

// 指纹去重
func (ec *Consume) removeDuplicates(alerts []models.AlertCurEvent) []models.AlertCurEvent {
	/*
		alert中有不重复字段，last_eval_time。
	*/

	latestAlert := make(map[string]models.AlertCurEvent)
	var newAlerts []models.AlertCurEvent

	for _, alert := range alerts {
		// 以最新为准
		latestAlert[alert.Fingerprint] = alert
	}

	for _, alert := range latestAlert {
		newAlerts = append(newAlerts, alert)
	}

	return newAlerts
}

// 触发告警通知
func (ec *Consume) fireAlertEvent(alertsMap map[string][]models.AlertCurEvent) {
	var wg sync.WaitGroup

	for _, alerts := range alertsMap {
		for _, alert := range alerts {
			alertCopy := alert
			wg.Add(1)
			go func(models.AlertCurEvent) {
				defer wg.Done()
				ec.addAlertToGroup(alertCopy)
				if alertCopy.IsRecovered {
					ec.removeAlertFromCache(alertCopy)
					//记录历史告警

					err := process.RecordAlertHisEvent(ec.ctx, alertCopy)
					if err != nil {
						global.Logger.Sugar().Error(err.Error())
						return
					}

				}
			}(alertCopy)
		}
	}

	wg.Wait()

	for _, alerts := range ec.preStoreAlertGroup {
		ec.handleAlert(alerts)
	}
}

// 删除缓存
func (ec *Consume) removeAlertFromCache(alert models.AlertCurEvent) {
	key := alert.GetFiringAlertCacheKey()
	ec.ctx.Redis.Event().DelCache(key)
}

// 添加告警到组(分组)
func (ec *Consume) addAlertToGroup(alert models.AlertCurEvent) {
	// 如果没有定义通知组，则直接添加到 ruleId 组中
	if alert.NoticeGroup == nil || len(alert.NoticeGroup) == 0 {
		ec.addAlertToGroupByRuleId(alert)
		return
	}

	// 遍历所有的 Metric
	matched := false
	for key, value := range alert.Metric {
		// 遍历所有的通知组
		for _, noticeGroup := range alert.NoticeGroup {
			// 如果当前 Metric 的 key 和 value 与通知组中的相匹配
			if noticeGroup["key"] == key && noticeGroup["value"] == value.(string) {
				// 计算分组的 ID 并添加警报到对应的组
				groupId := ec.calculateGroupHash(key, value.(string))
				ec.addAlertToGroupByGroupId(groupId, alert)
				matched = true
				break
			}
		}
		if matched {
			break
		}
	}

	// 如果没有找到任何匹配的组，则添加到 ruleId 组中
	if !matched {
		ec.addAlertToGroupByRuleId(alert)
	}
}

// 以Id作为key添加到组
func (ec *Consume) addAlertToGroupByGroupId(groupId string, alert models.AlertCurEvent) {
	ec.Lock()
	defer ec.Unlock()

	// 将告警和恢复消息再分组
	if alert.IsRecovered {
		groupId = "recovered-" + groupId
	}

	ec.preStoreAlertGroup[groupId] = append(ec.preStoreAlertGroup[groupId], alert)
}

// 以ruleName作为key添加到组
func (ec *Consume) addAlertToGroupByRuleId(alert models.AlertCurEvent) {
	ec.Lock()
	defer ec.Unlock()

	// 将告警和恢复消息再分组
	if alert.IsRecovered {
		alert.RuleId = "recovered-" + alert.RuleId
	}
	ec.preStoreAlertGroup[alert.RuleId] = append(ec.preStoreAlertGroup[alert.RuleId], alert)
}

// hash
func (ec *Consume) calculateGroupHash(key, value string) string {
	return hash.Md5Hash([]byte(key + ":" + value))
}

// 推送告警
func (ec *Consume) handleAlert(alerts []models.AlertCurEvent) {
	if alerts == nil {
		return
	}

	var (
		content = new(common.MyString)
		curTime = time.Now().Unix()
	)

	//通知模版 目前一条告警规则只能匹配一条告警模版
	noticeData, _ := ec.ctx.DB.Notice().Get(models.NoticeQuery{
		TenantId: alerts[0].TenantId,
		ID:       alerts[0].NoticeId,
	})

	if common.IsEmptyStr(noticeData.Name) {
		//通知模版不存在 返回错误信息
		return
	}

	for i := range alerts {
		//每次都需要情况
		var fetchedUsers []system.SysUser
		switch noticeData.UserNotices.ReceiverType {
		case "User":
			fetchedUsers = process.GetAlertUsers(ec.ctx, noticeData)
		case "Duty":
			//获取值班表的用户
			fetchedUser := process.GetDutyUser(ec.ctx, noticeData)
			fetchedUsers = append(fetchedUsers, *fetchedUser)
		}
		//更新告警人
		alerts[i].DutyUser = fetchedUsers
		//告警详情组合
		content.A(templates.DetailTemplate(alerts[i])).A("\n")

	}
	// 告警聚合,减少告警噪音， 每组告警取第一位的告警数据
	alerts[0].Annotations = content.Str()

	var wg sync.WaitGroup
	for i := range alerts {
		wg.Add(1)
		n := i
		go func(int) {
			defer wg.Done()
			if !alerts[n].IsRecovered {
				alerts[n].LastSendTime = curTime
				alerts[n].Duration = curTime - alerts[n].FirstTriggerTime
				ec.ctx.Redis.Event().SetCache("Firing", alerts[n], 0)
			}

		}(n)
	}
	wg.Wait()

	// 相同告警规则聚合到第一条告警
	err := sender.Sender(ec.ctx, alerts, noticeData)
	if err != nil {
		global.Logger.Sugar().Errorf(err.Error())
		return
	}

}
