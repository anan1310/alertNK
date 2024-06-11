package alert

import (
	"alarm_collector/alert/eval"
	"alarm_collector/pkg/ctx"
)

func Initialize(ctx *ctx.Context) {

	eval.NewInterAlertRuleWork(ctx).Run()

}
