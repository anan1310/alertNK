package alert

import (
	"alarm_collector/alert/consumer"
	"alarm_collector/alert/eval"
	"alarm_collector/pkg/ctx"
)

func Initialize(ctx *ctx.Context) {
	// 消费者
	consumer.NewInterEvalConsumeWork(ctx).Run()
	// 生产者
	eval.NewInterAlertRuleWork(ctx).Run()

}
