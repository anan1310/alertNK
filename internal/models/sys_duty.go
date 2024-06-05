package models

import "alarm_collector/pkg/utils/common"

// DutyManagement 值班日程
type DutyManagement struct {
	TenantId    string `json:"tenantId"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Manager     Users  `json:"manager" gorm:"manager;serializer:json"`
	Description string `json:"description"`
	CreateBy    string `json:"create_by"`
	CreateAt    int64  `json:"create_at"`
}

func (DutyManagement) TableName() string {
	return "sys_duty_management"
}

type DutyManagementQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	common.PageInfo
}
