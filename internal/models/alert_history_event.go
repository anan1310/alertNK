package models

// AlertHisEvent 历史告警
type AlertHisEvent struct {
	TenantId         string                 `json:"tenantId"`
	DatasourceId     string                 `json:"datasource_id" gorm:"datasource_id"`
	DatasourceType   string                 `json:"datasource_type"`
	Fingerprint      string                 `json:"fingerprint"`
	RuleId           string                 `json:"rule_id"`
	RuleName         string                 `json:"rule_name"`
	Severity         string                 `json:"severity"`
	Metric           map[string]interface{} `json:"metric" gorm:"metric;serializer:json"`
	EvalInterval     int64                  `json:"eval_interval"`
	Annotations      string                 `json:"annotations"`
	IsRecovered      bool                   `json:"is_recovered" gorm:"-"`
	FirstTriggerTime int64                  `json:"first_trigger_time"` // 第一次触发时间
	LastEvalTime     int64                  `json:"last_eval_time"`     // 最近评估时间
	LastSendTime     int64                  `json:"last_send_time"`     // 最近发送时间
	RecoverTime      int64                  `json:"recover_time"`       // 恢复时间
}

type AlertHisEventQuery struct {
	TenantId       string `json:"tenantId" form:"tenantId"`
	DatasourceId   string `json:"datasourceId" form:"datasourceId"`
	DatasourceType string `json:"datasourceType" form:"datasourceType"`
	Fingerprint    string `json:"fingerprint" form:"fingerprint"`
	Severity       string `json:"severity" form:"severity"`
	RuleId         string `json:"ruleId" form:"ruleId"`
	RuleName       string `json:"ruleName" form:"ruleName"`
	StartAt        int64  `json:"startAt" form:"startAt"`
	EndAt          int64  `json:"endAt" form:"endAt"`
	PageIndex      int64  `json:"pageIndex" form:"pageIndex"`
	PageSize       int64  `json:"pageSize" form:"pageSize"`
}

type HistoryEventResponse struct {
	List       []AlertHisEvent
	PageIndex  int64
	PageSize   int64
	TotalCount int64
}
