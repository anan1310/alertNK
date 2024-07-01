package templates

import (
	"alarm_collector/internal/models"
)

type Template struct {
	alerts []models.AlertCurEvent
	notice models.AlertNotice
}
type InterTemplate interface {
	SendAlertEmail() error
	SendAlertSMS() error
	SendAlertDingDing() error
}

func NewTemplate(alerts []models.AlertCurEvent, notice models.AlertNotice) InterTemplate {
	return Template{
		alerts: alerts,
		notice: notice,
	}
}
