package models

import (
	"gorm.io/gorm"
)

type AlertNotice struct {
	TenantId        string `json:"tenantId"`
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	Env             string `json:"env"`
	DutyId          string `json:"dutyId"`
	NoticeType      string `json:"noticeType"`
	EnableCard      string `json:"enableCard"`
	Hook            string `json:"hook"`
	Template        string `json:"template"`
	TemplateFiring  string `json:"templateFiring"`
	TemplateRecover string `json:"templateRecover"`
}

type AlertRecord struct {
	gorm.Model
	AlertName   string `json:"alertName"`
	Description string `json:"description"`
	Metric      string `json:"metric"`
	Severity    string `json:"severity"`
	Status      string `json:"status"`
}
