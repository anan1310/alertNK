package models

import "alarm_collector/pkg/utils/common"

type AlertNotice struct {
	TenantId        string `json:"tenantId"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	DutyId          string `json:"dutyId"`
	NoticeType      string `json:"noticeType"`
	EnableCard      string `json:"enableCard"`
	Hook            string `json:"hook"`
	Template        string `json:"template"`
	TemplateFiring  string `json:"templateFiring"`
	TemplateRecover string `json:"templateRecover"`
}

func (AlertNotice) TableName() string {
	return "sys_alert_notice"
}

type NoticeQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Query    string `json:"query" form:"query"`
	common.PageInfo
}
