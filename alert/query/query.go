package query

import (
	"alarm_collector/alert/process"
	"alarm_collector/alert/queue"
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/templates"
	"time"
)

type RuleQuery struct {
	alertEvent models.AlertCurEvent
	ctx        *ctx.Context
}

func (rq *RuleQuery) Query(ctx *ctx.Context, rule models.AlertRule) {
	rq.ctx = ctx

	switch rule.DatasourceType {
	case "Prometheus":
		rq.prometheus(rule)
	}

}

// 告警恢复
func (rq *RuleQuery) alertRecover(rule models.AlertRule, curKeys []string) {
	firingKeys, err := rq.ctx.Redis.Rule().GetAlertFiringCacheKeys(models.AlertRuleQuery{
		TenantId:         rule.TenantId,
		RuleId:           rule.RuleId,
		DatasourceIdList: rule.DatasourceIdList,
	})
	if err != nil {
		return
	}
	// 获取已恢复告警的keys
	recoverKeys := process.GetSliceDifference(firingKeys, curKeys)
	if recoverKeys == nil || len(recoverKeys) == 0 {
		return
	}

	curTime := time.Now().Unix()
	for _, key := range recoverKeys {
		event := rq.ctx.Redis.Event().GetCache(key)
		if event.IsRecovered == true {
			return
		}

		if _, exists := queue.RecoverWaitMap[key]; !exists {
			// 如果没有，则记录当前时间
			queue.RecoverWaitMap[key] = curTime
			continue
		}

		// 判断是否在等待时间范围内
		rt := time.Unix(queue.RecoverWaitMap[key], 0).Add(time.Minute * time.Duration(global.Config.Server.RecoverWait)).Unix()
		if rt > curTime {
			continue
		}

		event.IsRecovered = true
		event.RecoverTime = curTime
		event.LastSendTime = 0
		//处理告警内容
		event.Annotations = templates.DetailTemplate(event)
		rq.ctx.Redis.Event().SetCache("Firing", event, 0)

		// 触发恢复删除带恢复中的 key
		delete(queue.RecoverWaitMap, key)
	}
}

/*
	恢复告警逻辑很简单，alert/query/query.go(alertRecover 方法)，
	根据每个规则的query的response 和 redis缓存进行对比取差异值，
	（缓存中存在，query response不存在则视为恢复，会把恢复的key丢到 recoverWaitGroup中 等待多久后依然触发恢复则视为恢复，等待时间由配置文件的recoverWait决定）
*/

// Prometheus 数据源
func (rq *RuleQuery) prometheus(rule models.AlertRule) {
	//聚合告警规则
	var (
		curFiringKeys  = &[]string{}
		curPendingKeys = &[]string{}
	)
	//当前prometheus 执行完成后执行下面方法
	defer func() {
		go process.GcPendingCache(rq.ctx, rule, *curPendingKeys)
		rq.alertRecover(rule, *curFiringKeys)
		go process.GcRecoverWaitCache(rule, *curFiringKeys)
	}()

	//获取告警源数据
	alertSource, alarmRule, err := rq.handleAlertSource(rule)
	if err != nil {
		return
	}
	global.Logger.Sugar().Info("告警源前数据-->", alertSource)
	conditionStack, severity, err := evaluateRPN(alarmRule, alertSource)
	global.Logger.Sugar().Info("告警源后数据-->", alertSource)
	if err != nil {
		global.Logger.Sugar().Error("告警规则解析错误-->", err)
		return
	}

	// 如果最终条件为真，推送告警到redis中
	if conditionStack && len(alertSource) > 0 {
		process.CalIndicatorValue(rq.ctx, curFiringKeys, curPendingKeys, alertSource, rule, severity)
		global.Logger.Sugar().Info("%s:触发告警,告警规则", rule.RuleName)
	}

}
