package templates

import "alarm_collector/internal/models"

type Template struct {
	alert  models.AlertCurEvent
	notice models.AlertNotice
}
type InterTemplate interface {
	SendAlertEmail() error
	SendAlertSMS() error
	SendAlertDingDing() error
}

func NewTemplate(alert models.AlertCurEvent, notice models.AlertNotice) InterTemplate {
	return Template{
		alert:  alert,
		notice: notice,
	}
}
