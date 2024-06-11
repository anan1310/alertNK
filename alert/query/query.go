package query

import (
	"alarm_collector/alert/process"
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/common"
	"fmt"
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

// Prometheus 数据源
func (rq *RuleQuery) prometheus(rule models.AlertRule) {
	//聚合告警规则
	var (
		alarmRule     = new(common.MyString)
		rules         = rule.PrometheusConfig.Rules
		targetMapping = new(common.MyString)

		curFiringKeys  = &[]string{}
		curPendingKeys = &[]string{}
	)

	size := len(rules)
	for i, r := range rules {
		switch rule.PrometheusConfig.IsUnionRule {
		case 0:
			//system_cpu_usage < 0.1,||,load1 > 0.2
			alarmRule.A(r.TargetMapping + " " + r.Operator + " " + common.StrVal(r.Value) + " " + r.Severity)
			//拼接字段
			targetMapping.A(r.TargetMapping)
			if i < size-1 {
				targetMapping.A(",")
				alarmRule.A(",||,")
			}
		case 1:
		case 2:

		}
	}
	//获取数据源的值  如果达到告警的阈值 那么就写入redis缓冲中
	s := models.PrometheusDataSourceQuery{
		MetricType:    "linux",
		MetricName:    "master",
		TargetMapping: targetMapping.Str(),
		Pid:           "4",
	}
	//获取告警源
	alertSource, err := rq.ctx.CK.PrometheusDataSource().Get(s)
	if err != nil {
		return
	}
	global.Logger.Sugar().Info("告警源前数据-->", alertSource)
	err, conditionStack, severity := ParsePromRule(alarmRule.Str(), alertSource)
	global.Logger.Sugar().Info("告警源后数据-->", alertSource)
	fmt.Println(*severity)
	// 如果最终条件为真，则触发告警
	if len(conditionStack) == 1 && conditionStack[0] {
		//将告警推送到redis中
		process.CalIndicatorValue(rq.ctx, curFiringKeys, curPendingKeys, alertSource, rule, *severity)
		global.Logger.Sugar().Info("触发告警,告警规则")
	} else {
		return
	}

}
