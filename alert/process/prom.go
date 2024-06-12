package process

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/hash"
)

func CalIndicatorValue(ctx *ctx.Context, curFiringKeys, curPendingKeys *[]string, metricMap map[string]interface{}, rule models.AlertRule, severity string) {
	event := ParserDefaultEvent(rule)
	event.Fingerprint = hash.GenerateFingerprint(rule.RuleId) //告警指纹
	event.DatasourceId = rule.DatasourceIdList[0]
	event.Metric = metricMap
	event.Metric["severity"] = severity
	event.Severity = severity
	//触发的告警规则
	for _, rules := range rule.PrometheusConfig.Rules {
		if _, ok := metricMap[rules.TargetMapping]; ok {
			event.Rules = append(event.Rules, rules)
		}
	}
	firingKey := event.GetFiringAlertCacheKey()
	pendingKey := event.GetPendingAlertCacheKey()

	*curFiringKeys = append(*curFiringKeys, firingKey)
	*curPendingKeys = append(*curPendingKeys, pendingKey)

	ok := ctx.DB.Rule().GetRuleIsExist(event.RuleId)
	if ok {
		SaveEventCache(ctx, event)
	}
}
