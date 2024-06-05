package models

type DutyScheduleCreate struct {
	TenantId   string  `json:"tenantId"`
	DutyId     string  `json:"dutyId"`
	DutyPeriod int     `json:"dutyPeriod"`
	Month      string  `json:"month"`
	Users      []Users `json:"users"`
}

type DutySchedule struct {
	TenantId string `json:"tenantId"`
	DutyId   string `json:"dutyId"`
	Time     string `json:"time"`
	Users
}

type Users struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
}

type DutyScheduleQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	DutyId   string `json:"dutyId" form:"dutyId"`
	Time     string `json:"time" form:"time"`
}

func (DutySchedule) TableName() string {
	return "sys_duty_schedule"
}
