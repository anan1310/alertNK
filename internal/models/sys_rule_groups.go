package models

import "alarm_collector/pkg/utils/common"

// RuleGroups 告警组
type RuleGroups struct {
	TenantId    string `json:"tenantId"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (RuleGroups) TableName() string {
	return "sys_rule_groups"
}

type RuleGroupQuery struct {
	TenantId    string `json:"tenantId" form:"tenantId"`
	ID          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	common.PageInfo
}
