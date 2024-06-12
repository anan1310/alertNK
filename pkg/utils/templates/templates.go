package templates

import "alarm_collector/internal/models"

type Template struct {
	CardContentMsg string
}

func NewTemplate(alert models.AlertCurEvent, notice models.AlertNotice) Template {
	switch notice.NoticeType {
	case "FeiShu":
		return Template{CardContentMsg: "飞书"}
	case "DingDing":
		return Template{CardContentMsg: "钉钉"}
	case "QQ":
		return Template{}
	case "SMS":
		return Template{}
	}

	return Template{}
}
