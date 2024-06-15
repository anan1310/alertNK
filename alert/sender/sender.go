package sender

import (
	"alarm_collector/alert/mute"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/templates"
)

func Sender(ctx *ctx.Context, alert models.AlertCurEvent, notice models.AlertNotice) error {
	// 开启静默规则
	ok := mute.IsMuted(ctx, &alert)
	if ok {
		return nil
	}

	//获取告警信息
	interTemplate := templates.NewTemplate(alert, notice)

	switch notice.NoticeType {
	case "DingDing":
		err := interTemplate.SendAlertSMS()
		if err != nil {
			return err
		}
	}

	return nil
}
