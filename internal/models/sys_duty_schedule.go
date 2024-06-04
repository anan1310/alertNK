package models

type DutyManagement struct {
	TenantId    string `json:"tenantId"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Manager     Users  `json:"manager" gorm:"manager;serializer:json"`
	Description string `json:"description"`
	CreateBy    string `json:"create_by"`
	CreateAt    int64  `json:"create_at"`
}

type DutyManagementQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
}

type DutyScheduleCreate struct {
	TenantId   string  `json:"tenantId"`
	DutyId     string  `json:"dutyId"`
	DutyPeriod int     `json:"dutyPeriod"`
	Month      string  `json:"month"`
	Users      []Users `json:"users"`
}

type Users struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
}

type DutySchedule struct {
	TenantId string `json:"tenantId"`
	DutyId   string `json:"dutyId"`
	Time     string `json:"time"`
	Users
}

type DutyScheduleQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	DutyId   string `json:"dutyId" form:"dutyId"`
	Time     string `json:"time" form:"time"`
}
