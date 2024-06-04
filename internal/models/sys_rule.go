package models

type LabelsMap map[string]string

type NoticeGroup []map[string]string

type AlertRule struct {
	TenantId             string   `json:"tenantId"`
	RuleId               string   `json:"ruleId" gorm:"ruleId"`
	RuleGroupId          string   `json:"ruleGroupId"`                                      //所属告警组
	DatasourceType       string   `json:"datasourceType"`                                   //监控类型
	StrategyType         string   `json:"strategyType"`                                     //策略类型
	DatasourceIdList     []string `json:"datasourceId" gorm:"datasourceId;serializer:json"` //数据源列表
	RuleName             string   `json:"ruleName"`                                         //规则名称
	Severity             string   `json:"severity"`                                         //告警等级 提示 严重 紧急
	EvalInterval         int64    `json:"evalInterval"`                                     //执行频率
	ForDuration          int64    `json:"forDuration"`                                      //持续时间
	RepeatNoticeInterval int64    `json:"repeatNoticeInterval"`                             //重复通知间隔时间
	Description          string   `json:"description"`                                      //描述信息

	NoticeId    string      `json:"noticeId"`
	NoticeGroup NoticeGroup `json:"noticeGroup" gorm:"noticeGroup;serializer:json"`
	Trigger     RuleTrigger `json:"trigger" gorm:"foreignKey:RuleId;references:RuleId"` //关联触发条件
	Enabled     *bool       `json:"enabled" gorm:"enabled"`
}

func (AlertRule) TableName() string {
	return "sys_alert_rules"
}

type AlertRuleQuery struct {
	TenantId         string   `json:"tenantId" form:"tenantId"`
	RuleId           string   `json:"ruleId" form:"ruleId"`
	RuleGroupId      string   `json:"ruleGroupId" form:"ruleGroupId"`
	DatasourceType   string   `json:"datasourceType" form:"datasourceType"`
	DatasourceIdList []string `json:"datasourceId" form:"datasourceId"`
	RuleName         string   `json:"ruleName" form:"ruleName"`
	Enabled          string   `json:"enabled" form:"enabled"`
}
