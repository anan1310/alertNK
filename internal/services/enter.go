package services

import (
	"alarm_collector/pkg/ctx"
)

var (
	UserService         interSysUserService
	AlertService        interAlertService //告警规则推送
	RuleGroupService    interRuleGroupService
	DutyManagerService  interDutyManagerService
	DutyCalendarService interDutyCalendarService
	NoticeService       interNoticeService
	SilenceService      interSilenceService
)

func NewServices(ctx *ctx.Context) {
	UserService = newInterUserService(ctx)
	AlertService = newInterAlertService(ctx)
	RuleGroupService = newInterRuleGroupService(ctx)
	DutyManagerService = newInterDutyMangerService(ctx)
	DutyCalendarService = newInterDutyCalendarService(ctx)
	NoticeService = newInterAlertNoticeService(ctx)
	SilenceService = newInterSilenceService(ctx)
}
