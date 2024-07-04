package templates

import "alarm_collector/internal/models"

// 显示详情

func DetailTemplate(alert models.AlertCurEvent) string {

	templateStr := `
{{- define "Title" -}}
{{- if not .IsRecovered }}
[报警中] 🔥
{{- else }}
[已恢复] ✨
{{- end -}}
{{- end -}}

{{- define "TitleColor" -}}
{{- if not .IsRecovered -}}red{{- else -}}green{{- end -}}
{{- end -}}

{{ define "SeverityDescription" -}}
{{- if eq .Severity "P0" -}}紧急
{{- else if eq .Severity "P1" -}}严重
{{- else if eq .Severity "P2" -}}提示
{{- else -}}未知
{{- end -}}
{{ end -}}

{{- define "Event" -}}
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
📌 告警等级: {{ template "SeverityDescription" . }} 
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
{{- end -}}

{{- define "Footer" -}}
{{- end -}}
`

	Title := ParserTemplate("Title", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)

	t := Title + "\n" + Event + "\n"

	return t
}
