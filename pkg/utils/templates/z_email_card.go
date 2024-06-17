package templates

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/common"
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

	emailBody := new(common.MyString)
	// 生成邮件内容
	for i, alert := range t.alerts {
		template := emailTemplate(alert)
		if i < len(t.alerts) {
			emailBody.A(fmt.Sprintf("第 %d 告警规则信息：\n", i))
		}
		emailBody.A(template).A("\n")
	}

	// 创建新的邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", t.alerts[0].DutyUser.Email)
	m.SetHeader("Subject", "告警通知")
	m.SetBody("text/html", emailBody.Str())

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
        {{- if not .IsRecovered -}}
            【报警中】 🔥<br>
        {{- else -}}
            【已恢复】 ✨<br>
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
           <br>
            **🤖 告警类型:** ${rule_name}<br>
            **🫧 告警指纹:** ${fingerprint}<br>
            **📌 告警等级:** ${severity}<br>
            **🖥 告警主机:** ${metric.instance}<br>
            **🕘 开始时间:** ${first_trigger_time_format}<br>
            **👤 值班人员:** ${duty_user.user_name}<br>
            **📝 报警事件:** {{ range .Rules -}}
		              {{ .MetricName }} {{ .Operator }} {{ .Value }}{{ .Unit }}, <br>
		          {{- end }}
        {{- else -}}
            **🤖 告警类型:** ${rule_name}<br>
            **🫧 告警指纹:** ${fingerprint}<br>
            **📌 告警等级:** P${severity}<br>
            **🖥 告警主机:** ${metric.instance}<br>
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
