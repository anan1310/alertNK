package models

import "alarm_collector/pkg/utils/common"

const SilenceCachePrefix = "mute-"

type AlertSilences struct {
	TenantId       string `json:"tenantId"`
	Id             string `json:"id"`
	Fingerprint    string `json:"fingerprint"`
	Datasource     string `json:"datasource"`
	DatasourceType string `json:"datasource_type"`
	StartsAt       int64  `json:"starts_at"`
	EndsAt         int64  `json:"ends_at"`
	CreateBy       string `json:"create_by"`
	UpdateBy       string `json:"update_by"`
	CreateAt       int64  `json:"create_at"`
	UpdateAt       int64  `json:"update_at"`
	Comment        string `json:"comment"`
}

func (AlertSilences) TableName() string {
	return "sys_alert_silences"
}

type AlertSilenceQuery struct {
	TenantId       string `json:"tenantId" form:"tenantId"`
	ID             string `json:"id" form:"id"`
	Fingerprint    string `json:"fingerprint" form:"fingerprint"`
	Datasource     string `json:"datasource" form:"datasource"`
	DatasourceType string `json:"datasourceType" form:"datasourceType"`
	Comment        string `json:"comment" form:"comment"`
	common.PageInfo
}
