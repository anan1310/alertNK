package models

// RuleTrigger 告警条件规则
type RuleTrigger struct {
	TriggerId   int     `gorm:"column:trigger_id" json:"triggerId"`
	MetricName  string  `gorm:"column:metric_name" json:"metricName"`  //指标名称
	Unit        string  `gorm:"column:unit" json:"unit"`               // 告警指标单位
	Value       float64 `gorm:"column:value" json:"value"`             // 告警指标值
	Operator    string  `gorm:"column:operator" json:"operator"`       // 告警操作符
	Description string  `gorm:"column:description" json:"description"` // 告警描述
	RuleId      string  `gorm:"column:rule_id" json:"RuleId"`          //关联告警
}

func (s *RuleTrigger) TableName() string {
	return "sys_rule_trigger"
}
