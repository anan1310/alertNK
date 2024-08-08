package eval

import (
	"alarm_collector/alert/query"
	"alarm_collector/alert/queue"
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/ctx"
	"context"
	"sync"
	"time"
)

type AlertRuleWork struct {
	sync.RWMutex
	query.RuleQuery
	ctx        *ctx.Context
	rule       chan *models.AlertRule
	alertEvent models.AlertCurEvent
}

type InterAlertRuleWork interface {
	Run()
}

func NewInterAlertRuleWork(ctx *ctx.Context) InterAlertRuleWork {
	return &AlertRuleWork{
		rule: queue.AlertRuleChannel,
		ctx:  ctx,
	}
}

// Run 持续获取告警规则的状态
func (arw *AlertRuleWork) Run() {

	go func() {
		for {
			select {
			case rule := <-arw.rule:
				if *rule.Enabled {
					// 创建一个用于停止协程的上下文
					c, cancel := context.WithCancel(context.Background())
					queue.WatchCtxMap[rule.RuleId] = cancel
					go arw.worker(*rule, c)
				}
			}
		}
	}()

	// 重启后 将历史规则重新推送到 arw.rule 通道中。
	services.AlertService.RePushRule(arw.ctx, arw.rule)

}

func (arw *AlertRuleWork) workerBak(rule models.AlertRule, ctx context.Context) {
	// 执行频率 比如10秒一次
	ei := time.Second * time.Duration(rule.EvalInterval)
	timer := time.NewTimer(ei)

	for {
		select {
		case <-timer.C:
			global.Logger.Sugar().Infof("规则评估 -> %v", rule)
			arw.Query(arw.ctx, rule)

		case <-ctx.Done():
			global.Logger.Sugar().Infof("停止 RuleId 为 %v 的 Watch 协程", rule.RuleId)
			return
		}
		// 重置执行频率
		timer.Reset(time.Second * time.Duration(rule.EvalInterval))

	}

}

func (arw *AlertRuleWork) worker(rule models.AlertRule, ctx context.Context) {
	// 执行频率，比如10秒一次
	ei := time.Second * time.Duration(rule.EvalInterval)
	ticker := time.NewTicker(ei)
	defer ticker.Stop() // 在函数退出前停止ticker，释放资源

	for {
		select {
		case <-ticker.C:
			global.Logger.Sugar().Infof("规则评估 -> %v", rule)
			arw.Query(arw.ctx, rule)

		case <-ctx.Done():
			global.Logger.Sugar().Infof("停止 RuleId 为 %v 的 Watch 协程", rule.RuleId)
			return
		}
	}
}

/*
Run 方法主要功能是通过启动一个无限循环的 Goroutine 来持续监听和处理告警规则，并在服务重启后重新推送历史规则以保持规则的连续性和有效性。
*/
