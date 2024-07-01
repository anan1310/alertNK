package models

import "alarm_collector/pkg/utils/common"

type AlertNotice struct {
	TenantId             string      `json:"tenantId"`
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	EnabledAlertNotice   *bool       `json:"enabledAlertNotice"`   //告警通知
	EnabledRecoverNotice *bool       `json:"enabledRecoverNotice"` //告警恢复通知
	UserNotices          UserNotices `json:"userNotices"  gorm:"user_notices;serializer:json"`
}

type UserNotices struct {
	ReceiverType string   `json:"receiverType"`   //Duty(值班表)  User(用户)
	DutyId       string   `json:"dutyId"`         //值班表iD
	UserIds      []int    `json:"userIds"`        //用户ID
	NoticeWay    []string `json:"noticeWay"`      //通知渠道
	Week         []string `json:"week"`           //通知周期
	StartTime    int      `json:"startTime"`      //通知时段
	EndTime      int      `json:"endTime"`        // StartTime - EndTime
	Hook         string   `json:"hook,omitempty"` //回调函数
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
