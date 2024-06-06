package system

type RouterGroup struct {
	UserRouter
	RuleGroupRouter
	DutyManagerRouter
	DutyCalendarRouter
	NoticeRouter
	SilencesRouter
}
