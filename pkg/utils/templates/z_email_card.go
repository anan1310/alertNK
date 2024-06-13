package templates

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"fmt"
	"gopkg.in/gomail.v2"
)

//邮箱模版

func (t Template) SendAlertEmail() error {
	// 配置SMTP服务器
	smtpHost := global.Config.Mail.Host
	smtpPort := global.Config.Mail.Port
	smtpUser := global.Config.Mail.SmtpUser
	smtpPass := global.Config.Mail.Pass

	// 生成邮件内容
	emailBody := emailTemplate(t.alert)

	// 创建新的邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", t.alert.DutyUser.Email)
	m.SetHeader("Subject", "告警通知")
	m.SetBody("text/html", emailBody)

	// 发送邮件
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	fmt.Println(d)
	/*
		err := d.DialAndSend(m)
		if err != nil {
			return err
		}

	*/

	return nil
}

func emailTemplate(alert models.AlertCurEvent) string {
	templateStr := `
	{{- define "Title" -}}
        {{- if not .IsRecovered -}}
            【报警中】- 即时设计业务系统 🔥<br>
        {{- else -}}
            【已恢复】- 即时设计业务系统 ✨<br>
        {{- end -}}
    {{- end }}

    {{- define "TitleColor" -}}
        {{- if not .IsRecovered -}}
            red
        {{- else -}}
            green
        {{- end -}}
    {{- end }}

    {{ define "Event" -}}
        {{- if not .IsRecovered -}}
            **🤖 报警类型:** ${rule_name}<br>
            **🫧 报警指纹:** ${fingerprint}<br>
            **📌 报警等级:** ${severity}<br>
            **🖥 报警主机:** ${metric.instance}<br>
            **🕘 开始时间:** ${first_trigger_time_format}<br>
            **👤 值班人员:** ${duty_user.user_name}<br>
            **📝 报警事件:**
       {{- range .rules }}
		[${severity}] {{ .metricName }} {{ .operator }} {{ .value }}{{ .unit }}，统计粒度{{ .TargetExpression }}，连续1次满足条件则每1小时告警一次<br>
		{{- end }}
        {{- else -}}
            **🤖 报警类型:** ${rule_name}<br>
            **🫧 报警指纹:** ${fingerprint}<br>
            **📌 报警等级:** P${severity}<br>
            **🖥 报警主机:** ${metric.instance}<br>
            **🕘 开始时间:** ${first_trigger_time_format}<br>
            **🕘 恢复时间:** ${recover_time_format}<br>
            **👤 值班人员:** ${duty_user.user_name}<br>
            **📝 报警事件:** ${annotations}<br>
        {{- end -}}
    {{ end }}

    {{- define "Footer" -}}
        🧑‍💻 即时设计 - 运维团队
    {{- end }}
`

	Title := ParserTemplate("Title", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)
	Footer := ParserTemplate("Footer", alert, templateStr)

	t := Title + "\n" + Event + "\n" + Footer

	return t
}
