package sender

import (
	"alarm_collector/alert/mute"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/templates"
	"bytes"
	"fmt"
)

func Sender(ctx *ctx.Context, alert models.AlertCurEvent, notice models.AlertNotice) error {
	ok := mute.IsMuted(ctx, &alert)
	if ok {
		return nil
	}
	//获取告警信息
	n := templates.NewTemplate(alert, notice)

	cardContentByte := bytes.NewReader([]byte(n.CardContentMsg))
	fmt.Println(cardContentByte)

	return nil
}
