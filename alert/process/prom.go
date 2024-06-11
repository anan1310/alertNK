package process

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
)

func CalIndicatorValue(ctx *ctx.Context, curFiringKeys, curPendingKeys *[]string, metricMap map[string]interface{}, rule models.AlertRule, severity string) {
	event := ParserDefaultEvent(rule)
	//event.Fingerprint = v.GetFingerprint()
	event.Metric = metricMap
	event.Metric["severity"] = severity
	event.Severity = severity

	firingKey := event.GetFiringAlertCacheKey()
	pendingKey := event.GetPendingAlertCacheKey()

	*curFiringKeys = append(*curFiringKeys, firingKey)
	*curPendingKeys = append(*curPendingKeys, pendingKey)

	ok := ctx.DB.Rule().GetRuleIsExist(event.RuleId)
	if ok {
		SaveEventCache(ctx, event)
	}
}
