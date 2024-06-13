package templates

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/cmd"
	"bytes"
	"encoding/json"
	"text/template"
	"time"
)

var tmpl *template.Template

// ParserTemplate 处理告警推送的消息模版
func ParserTemplate(defineName string, alert models.AlertCurEvent, templateStr string) string {

	firstTriggerTime := time.Unix(alert.FirstTriggerTime, 0).Format(global.Layout)
	recoverTime := time.Unix(alert.RecoverTime, 0).Format(global.Layout)
	alert.FirstTriggerTimeFormat = firstTriggerTime
	alert.RecoverTimeFormat = recoverTime

	tmpl = template.Must(template.New("tmpl").Parse(templateStr))

	var (
		buf bytes.Buffer
		err error
	)

	err = tmpl.ExecuteTemplate(&buf, defineName, alert)
	if err != nil {
		global.Logger.Sugar().Error("告警模版执行失败 ->", err.Error())
		return ""
	}

	// 前面只会渲染出模版框架, 下面来渲染告警数据内容
	if defineName == "Event" {
		data := parserEvent(alert)
		return cmd.ParserVariables(buf.String(), data)
	}

	return buf.String()

}

func parserEvent(alert models.AlertCurEvent) map[string]interface{} {

	data := make(map[string]interface{})

	eventJson := cmd.JsonMarshal(alert)
	err := json.Unmarshal([]byte(eventJson), &data)
	if err != nil {
		global.Logger.Sugar().Error("parserEvent Unmarshal failed: ", err)
	}
	
	return data

}
