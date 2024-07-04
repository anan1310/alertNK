package sender

import (
	"alarm_collector/alert/mute"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"alarm_collector/pkg/utils/templates"
	"sync"
)

var wg sync.WaitGroup

func Sender(ctx *ctx.Context, alerts []models.AlertCurEvent, notice models.AlertNotice) error {
	// 开启静默规则

	ok := mute.IsMuted(ctx, &alerts[0], notice)
	if ok {
		return nil
	}

	//获取告警信息
	interTemplate := templates.NewTemplate(alerts, notice)
	wg.Add(len(notice.UserNotices.NoticeWay))
	// 启动协程发送
	for _, way := range notice.UserNotices.NoticeWay {
		noticeWay := way
		go func(noticeWay string) {
			defer wg.Done()
			switch noticeWay {
			case "DingDing":
				if err := interTemplate.SendAlertDingDing(); err != nil {
					return
				}
			case "Email":
				if err := interTemplate.SendAlertEmail(); err != nil {
					return
				}
			case "SMS":
				if err := interTemplate.SendAlertSMS(); err != nil {
					return
				}
			case "FeiShu":

			}
		}(noticeWay)
	}
	wg.Wait()

	return nil
}
