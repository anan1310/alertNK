package sender

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/templates"
)

func Sender(ctx *ctx.Context, alerts []models.AlertCurEvent, notice models.AlertNotice) error {
	// 开启静默规则
	/*
		ok := mute.IsMuted(ctx, &alerts[0])
		if ok {
			return nil
		}
	*/

	//获取告警信息
	interTemplate := templates.NewTemplate(alerts, notice)

	switch notice.NoticeType {
	case "DingDing":
		if err := interTemplate.SendAlertDingDing(); err != nil {
			return err
		}
	case "Email":
		if err := interTemplate.SendAlertEmail(); err != nil {
			return err
		}
	case "SMS":
		if err := interTemplate.SendAlertSMS(); err != nil {
			return err
		}
	}

	return nil
}
