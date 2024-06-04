package services

import (
	"alarm_collector/pkg/ctx"
)

var (
	UserService      InterSysUserService
	AlertService     InterAlertService //告警规则推送
	RuleGroupService InterRuleGroupService
)

func NewServices(ctx *ctx.Context) {
	UserService = newInterUserService(ctx)
	AlertService = newInterAlertService(ctx)
	RuleGroupService = newInterRuleGroupService(ctx)
}
