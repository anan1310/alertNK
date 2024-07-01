package templates

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/cmd"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/http_util"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Sms struct {
	Account   string //账户
	Timestamp int64  //时间戳（毫秒级别）
	Nonce     int    //随机正整数，参与加密
	Signature string //明文 填写原始密码
	Params    string //
}

func (t Template) SendAlertSMS() error {

	var (
		account      = "410710" //账户
		nonce        = 502175   //随机整数，默认和初始密码相同
		signature    = "502175" //密码（未加密）
		smsSignature = "【盛易信达】" //短信签名
	)
	content := new(common.MyString)

	var phoneNumbers []string
	for _, u := range t.alerts[0].DutyUser {
		phoneNumbers = append(phoneNumbers, u.PhoneNumber)
	}
	phoneNumber := strings.Join(phoneNumbers, ",")
	//短信告警人
	if common.IsEmptyStr(phoneNumber) {
		return fmt.Errorf("无效的手机号码")
	}
	//短信内容
	content.A(smsSignature).A(t.alerts[0].Annotations)
	//短信内容
	/*
		for i, alert := range t.alerts {
			smsContent := smsTemplate(alert)
			if i < len(t.alerts) {
				content.A(fmt.Sprintf("第 %d 告警规则信息：\n", i+1))
			}
			content.A(smsContent).A("\n")
		}
	*/
	//mobile 手机号，多个英文逗号分隔，最多1W个
	params := map[string]string{
		"mobile":  phoneNumber,
		"content": content.Str(),
	}
	marshal, _ := json.Marshal(params)
	sms := Sms{
		Account:   account,
		Timestamp: common.TimeUnixNano(),
		Nonce:     nonce,
		Signature: signature,
		Params:    string(marshal),
	}
	//发送告警短信
	err := sms.sendSms()
	if err != nil {
		return fmt.Errorf("短信告警错误:%v", err)
	}

	return nil
}

func (s *Sms) sendSms() error {

	cardContentString := cmd.JsonMarshal(s)
	smsContent := bytes.NewReader([]byte(cardContentString))

	_, err := http_util.Post("http://114.118.2.242:8090/sms/v1/message/send", smsContent)
	if err != nil {
		return err
	}

	return err
}

func smsTemplate(alert models.AlertCurEvent) string {
	// 定义模板字符串
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
			{{.MetricName}} {{.Operator}} {{.Value}}{{.Unit}}, 
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
			{{.MetricName}} {{.Operator}} {{.Value}}{{.Unit}}, 
			{{- end }}
	{{- end -}}
	{{ end }}
	
	{{- define "Footer" -}}
	{{- end }}
	`

	Title := ParserTemplate("Title", alert, templateStr)
	Event := ParserTemplate("Event", alert, templateStr)
	//Footer := ParserTemplate("Footer", alert, templateStr)

	t := Title + "\n" + Event + "\n"

	return t
}
