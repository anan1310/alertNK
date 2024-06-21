package system

type RouterGroup struct {
	UserRouter
	RuleGroupRouter
	RuleRouter
	DutyManagerRouter
	DutyCalendarRouter
	NoticeRouter
	SilencesRouter
	AlertEventRouter
	DashBoardInfoRouter
}
