package queue

import (
	"alarm_collector/internal/models"
	"context"
)

var (
	// WatchCtxMap 用于存储每个协程的上下文
	WatchCtxMap = make(map[string]context.CancelFunc)

	// AlertRuleChannel 用于消费用户创建的 Rule
	AlertRuleChannel = make(chan *models.AlertRule)

	// RecoverWaitMap 存储等待被恢复的告警的 Key
	RecoverWaitMap = make(map[string]int64)
)

/*
	WatchCtxMap：
	这通常用于在监控规则被禁用时，取消该规则对应的监控任务，以避免继续执行不必要的操作或占用资源。
	cancel()
	1:终止正在进行的操作
	2:释放资源
	3:通知所有关联的子上下文和操作停止执行
*/
