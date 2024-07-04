package templates

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/common"
	"gopkg.in/gomail.v2"
)

//邮箱模版

func (t Template) SendAlertEmail() error {
	// 配置SMTP服务器
	smtpHost := global.Config.Mail.Host
	smtpPort := global.Config.Mail.Port
	smtpUser := global.Config.Mail.SmtpUser
	smtpPass := global.Config.Mail.Pass

	emailBody := new(common.MyString)
	// 生成邮件内容
	/*
		for i, alert := range t.alerts {
			template := emailTemplate(alert)
			if i < len(t.alerts) {
				emailBody.A(fmt.Sprintf("第 %d 告警规则信息：\n", i+1))
			}
			emailBody.A(template).A("\n")
		}
	*/
	emailBody.A(t.alerts[0].Annotations)
	user := t.alerts[0].DutyUser
	var to []string
	for _, u := range user {
		to = append(to, u.Email)
	}
	// 创建新的邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "告警通知")
	//m.SetBody("text/html", emailBody.Str())
	m.SetBody("text/plain", emailBody.Str())

	// 发送邮件
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

func emailTemplate(alert models.AlertCurEvent) string {
	templateStr := `
	{{- define "Title" -}}
	{{- if not .IsRecovered -}}[报警中] 🔥{{- else -}}[已恢复] ✨{{- end -}}
	{{- end }}
	
	{{- define "TitleColor" -}}
	{{- if not .IsRecovered -}}red{{- else -}}green{{- end -}}
	{{- end }}
	
	{{ define "SeverityDescription" -}}
	{{- if eq .Severity "P0" }}紧急
	{{- else if eq .Severity "P1" }}严重
	{{- else if eq .Severity "P2" }}提示
	{{- else }}未知
	{{- end }}
	{{ end }}
	
	{{ define "Event" -}}
	{{- if not .IsRecovered -}}
	🤖 告警类型: {{.RuleName}}
	🫧 告警指纹: {{.Fingerprint}}
	📌 告警等级: {{ template "SeverityDescription" . }}
	🖥 告警主机: {{ .Metric.instance }}
	🕘 开始时间: {{.FirstTriggerTimeFormat}}
	👤 值班人员: {{ range .DutyUser -}}
			   {{.UserName}},
			   {{- end }}
	📝 报警事件: {{ range .Rules -}}
			{{.MetricName}} {{.Operator}} {{.Value}}{{.ToUnit}}, 
			{{- end }}
	{{- else -}}
	🤖 告警类型: {{.RuleName}}
	🫧 告警指纹: {{.Fingerprint}}
	📌 告警等级: {{.Severity}}
	🖥 告警主机: {{ .Metric.instance }}
	🕘 开始时间: {{.FirstTriggerTimeFormat}}
	🕘 恢复时间: {{.RecoverTimeFormat}}
	👤 值班人员: {{ range .DutyUser -}}
			   {{.UserName}},
			   {{- end }}
	📝 报警事件: {{ range .Rules -}}
			{{.MetricName}} {{.Operator}} {{.Value}}{{.ToUnit}}, 
			{{- end }}
	{{- end -}}
	{{ end }}
	
	{{- define "Footer" -}}
	{{- end }}
	`

	Title := ParserTemplate("Title", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)

	t := Title + "\n" + Event + "\n"

	return t
}
