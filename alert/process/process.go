package process

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
)

// GetRedisFiringKeys 获取缓存所有Firing的Keys
func GetRedisFiringKeys(ctx *ctx.Context) []string {
	var keys []string
	cursor := uint64(0)
	pattern := "*" + ":" + models.FiringAlertCachePrefix + "*"
	// 每次获取的键数量
	count := int64(100)

	for {
		var curKeys []string
		var err error

		curKeys, cursor, err = ctx.Redis.Redis().Scan(cursor, pattern, count).Result()
		if err != nil {
			break
		}

		keys = append(keys, curKeys...)

		if cursor == 0 {
			break
		}
	}

	return keys
}

// ParserDefaultEvent 解析默认告警事件
func ParserDefaultEvent(rule models.AlertRule) models.AlertCurEvent {

	event := models.AlertCurEvent{
		TenantId:             rule.TenantId,
		DatasourceType:       rule.DatasourceType,
		RuleId:               rule.RuleId,
		RuleName:             rule.RuleName,
		EvalInterval:         rule.EvalInterval,
		ForDuration:          rule.PrometheusConfig.ForDuration,
		NoticeId:             rule.NoticeId,
		NoticeGroup:          rule.NoticeGroup,
		IsRecovered:          false,
		RepeatNoticeInterval: rule.RepeatNoticeInterval,
		DutyUser:             "暂无", // 默认暂无值班人员, 渲染模版时会实际判断 Notice 是否存在值班人员
		Severity:             rule.Severity,
		EffectiveTime:        rule.EffectiveTime,
	}

	return event

}

func SaveEventCache(ctx *ctx.Context, event models.AlertCurEvent) {
	ctx.Lock()
	defer ctx.Unlock()

	firingKey := event.GetFiringAlertCacheKey()
	pendingKey := event.GetPendingAlertCacheKey()

	// 判断改事件是否是Firing状态, 如果不是Firing状态 则标记Pending状态
	resFiring := ctx.Redis.Event().GetCache(firingKey)
	if resFiring.Fingerprint != "" {
		event.FirstTriggerTime = resFiring.FirstTriggerTime
		event.LastEvalTime = ctx.Redis.Event().GetLastEvalTime(firingKey)
		event.LastSendTime = resFiring.LastSendTime
	} else {
		event.FirstTriggerTime = ctx.Redis.Event().GetFirstTime(pendingKey)
		event.LastEvalTime = ctx.Redis.Event().GetLastEvalTime(pendingKey)
		event.LastSendTime = ctx.Redis.Event().GetLastSendTime(pendingKey)
		ctx.Redis.Event().SetCache("Pending", event, 0)
	}

	// 初次告警需要比对持续时间
	if resFiring.LastSendTime == 0 {
		if event.LastEvalTime-event.FirstTriggerTime < event.ForDuration {
			return
		}
	}

	ctx.Redis.Event().SetCache("Firing", event, 0)
	ctx.Redis.Event().DelCache(pendingKey)

}
