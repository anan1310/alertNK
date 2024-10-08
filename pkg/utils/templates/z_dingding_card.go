package templates

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/models/system"
	"alarm_collector/pkg/utils/cmd"
	"alarm_collector/pkg/utils/http_util"
	"bytes"
	"fmt"
)

func (t Template) SendAlertDingDing() error {

	dingTemplate := bytes.NewReader([]byte(dingDingTemplate(t.alerts[0])))
	//dingTemplate := dingDingTemplate(t.alert)
	_, err := http_util.Post(t.notice.UserNotices.Hook, dingTemplate)
	if err != nil {
		return err
	}
	return nil
}

func dingDingTemplate(alert models.AlertCurEvent) string {
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
            **🤖 告警类型:** ${rule_name}<br>
            **🫧 告警指纹:** ${fingerprint}<br>
            **📌 告警等级:** ${severity}<br>
            **🖥 告警主机:** ${metric.instance}<br>
            **🕘 开始时间:** ${first_trigger_time_format}<br>
            **👤 值班人员:** ${duty_user.user_name}<br>
            **📝 报警事件:** {{ range .Rules -}}
		              {{ .MetricName }} {{ .Operator }} {{ .Value }}{{ .ToUnit }}, <br>
		          {{- end }}
        {{- else -}}
            **🤖 告警类型:** ${rule_name}<br>
            **🫧 告警指纹:** ${fingerprint}<br>
            **📌 告警等级:** ${severity}<br>
            **🖥 告警主机:** ${metric.instance}<br>
            **🕘 开始时间:** ${first_trigger_time_format}<br>
            **🕘 恢复时间:** ${recover_time_format}<br>
            **👤 值班人员:** ${duty_user.user_name}<br>
            **📝 报警事件:** ${annotations}<br>
        {{- end -}}
    {{ end }}

    {{- define "Footer" -}}
        
    {{- end }}
`

	Title := ParserTemplate("Title", alert, templateStr)
	TitleColor := ParserTemplate("TitleColor", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)
	//Footer := ParserTemplate("Footer", alert, templateStr)
	markdownContent := fmt.Sprintf("<font color=\"%s\">**%s**</font>\n\n%s\n\n", TitleColor, Title, Event)

	t := system.DingMsg{
		Msgtype: "markdown",
		Markdown: system.Markdown{
			Title: Title,
			Text:  markdownContent,
		},
		At: system.At{
			AtUserIds: []string{}, //被@人的用户userId。
			IsAtAll:   false,      //是否@所有人。
		},
	}
	cardContentString := cmd.JsonMarshal(t)
	return cardContentString
}
