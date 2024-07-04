package templates

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/internal/models/system"
	"fmt"
	"os"
	"testing"
	"text/template"
	"time"
)

func TestName(t *testing.T) {
	alert := models.AlertCurEvent{
		TenantId:       "tid-co4iic3adq7a2jjeas90",
		RuleId:         "r-cpkkccbvq9ld6m81o92g",
		RuleName:       "安志杰测试",
		DatasourceType: "Prometheus",
		DatasourceId:   "lp4Vo05wmoWbJGat9Slm8",
		Fingerprint:    "2898275c731beebb",
		Severity:       "P2",
		Metric: map[string]interface{}{
			"load1":    0.1,
			"severity": "P2",
			"instance": "1.1.1.1",
		},
		Labels:                 nil,
		EvalInterval:           10,
		ForDuration:            20,
		NoticeId:               "n-cpjqp73vq9l6bp9trkqg",
		NoticeGroup:            nil,
		Annotations:            "",
		IsRecovered:            false,
		FirstTriggerTime:       1718258063,
		FirstTriggerTimeFormat: time.Unix(1718258063, 0).Format(global.Layout),
		RepeatNoticeInterval:   60,
		LastEvalTime:           1718258287,
		LastSendTime:           1718258289,
		RecoverTime:            0,
		RecoverTimeFormat:      "",
		DutyUser:               []system.SysUser{{UserName: "安志杰"}},
		Rules: []models.Rules{
			{

				TargetMapping:    "load1",
				TargetExpression: "",
				MetricName:       "1分钟平均负载",
				ToUnit:           "",
				Value:            0.02,
				Operator:         ">",
				Severity:         "一般",
				Description:      "1分钟平均负载",
			},
		},
	}

	templateStr := `
	{{- define "Title" -}}
		{{- if not .IsRecovered -}}
			【报警中】- 即时设计业务系统 🔥
		{{- else -}}
			【已恢复】- 即时设计业务系统 ✨
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
			**🤖 报警类型:** {{.RuleName}}
			**🫧 报警指纹:** {{.Fingerprint}}
			**📌 报警等级:** {{.Severity}}
			**🖥 报警主机:** {{index .Metric "instance"}}
			**🕘 开始时间:** {{.FirstTriggerTimeFormat}}
			**👤 值班人员:** {{.DutyUser.UserName}}
			**📝 报警事件:** {{ range .Rules -}}
							  	[{{ .Severity }}] {{ .MetricName }} {{ .Operator }} {{ .Value }}{{ .ToUnit }}，统计粒度{{ .TargetExpression }}，连续1次满足条件则每1小时告警一次
							{{ end -}}
		{{- else -}}
			**🤖 报警类型:** {{.RuleName}}
			**🫧 报警指纹:** {{.Fingerprint}}
			**📌 报警等级:** {{.Severity}}
			**🖥 报警主机:** {{index .Metric "instance"}}
			**🕘 开始时间:** {{.FirstTriggerTimeFormat}}
			**🕘 恢复时间:** {{.RecoverTimeFormat}}
			**👤 值班人员:** {{.DutyUser.UserName}}
			**📝 报警事件:** {{.Annotations}}
		{{- end -}}
	{{ end }}

	{{- define "Footer" -}}
		🧑‍💻 即时设计 - 运维团队
	{{- end }}
	`

	tmpl := template.Must(template.New("tmpl").Parse(templateStr))

	err := tmpl.ExecuteTemplate(os.Stdout, "Event", alert)
	if err != nil {
		fmt.Printf("告警模版执行失败 -> %v", err)
	}
}
