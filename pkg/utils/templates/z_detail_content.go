package templates

import "alarm_collector/internal/models"

// 显示详情

func DetailTemplate(alert models.AlertCurEvent) string {
	templateStr := `
	{{- define "Title" -}}
        {{- if not .IsRecovered -}}
            【报警中】 🔥<br>
        {{- else -}}
            【已恢复】 ✨<br>
        {{- end -}}
    {{- end }}

    {{ define "Event" -}}
        {{- if not .IsRecovered -}}
           <br>
            **🤖 告警类型:** ${rule_name}
            **🫧 告警指纹:** ${fingerprint}
            **📌 告警等级:** ${severity}
            **🖥 告警主机:** ${metric.instance}
            **🕘 开始时间:** ${first_trigger_time_format}
            **👤 值班人员:** ${duty_user.user_name
            **📝 报警事件:** {{ range .Rules -}}
		              {{ .MetricName }} {{ .Operator }} {{ .Value }}{{ .Unit }}, 
		          {{- end }}
        {{- else -}}
            **🤖 告警类型:** ${rule_name}
            **🫧 告警指纹:** ${fingerprint}
            **📌 告警等级:** ${severity}
            **🖥 告警主机:** ${metric.instance}
            **🕘 开始时间:** ${first_trigger_time_format}
            **🕘 恢复时间:** ${recover_time_format}
            **👤 值班人员:** ${duty_user.user_name}
            **📝 报警事件:** ${annotations}
        {{- end -}}
    {{ end }}

`

	Title := ParserTemplate("Title", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)

	t := Title + "\n" + Event + "\n"

	return t
}
