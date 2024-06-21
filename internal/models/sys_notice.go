package models

import "alarm_collector/pkg/utils/common"

type AlertNotice struct {
	TenantId        string `json:"tenantId"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	DutyId          string `json:"dutyId"`
	NoticeType      string `json:"noticeType"`
	EnableCard      string `json:"enableCard,omitempty"`
	Hook            string `json:"hook,omitempty"`
	Template        string `json:"template,omitempty"`
	TemplateFiring  string `json:"templateFiring,omitempty"`
	TemplateRecover string `json:"templateRecover,omitempty"`
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
