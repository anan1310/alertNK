package models

import "alarm_collector/pkg/utils/common"

type LabelsMap map[string]string

type NoticeGroup []map[string]string

type AlertRule struct {
	TenantId             string        `json:"tenantId"`
	RuleId               string        `json:"ruleId" gorm:"ruleId"`
	RuleGroupId          string        `json:"ruleGroupId"`                                        //所属告警组
	DatasourceType       string        `json:"datasourceType"`                                     //监控类型
	StrategyType         string        `json:"strategyType"`                                       //策略类型
	DatasourceIdList     []string      `json:"datasourceId" gorm:"datasourceId;serializer:json"`   //数据源列表
	RuleName             string        `json:"ruleName"`                                           //规则名称
	EvalInterval         int64         `json:"evalInterval"`                                       //执行频率
	RepeatNoticeInterval int64         `json:"repeatNoticeInterval"`                               //重复通知间隔时间
	Description          string        `json:"description"`                                        //描述信息
	EffectiveTime        EffectiveTime `json:"effectiveTime" gorm:"effectiveTime;serializer:json"` //告警周期
	Severity             string        `json:"severity"`
	// Prometheus
	PrometheusConfig PrometheusConfig `json:"prometheusConfig" gorm:"prometheusConfig;serializer:json"`

	NoticeId    string      `json:"noticeId"`
	NoticeGroup NoticeGroup `json:"noticeGroup" gorm:"noticeGroup;serializer:json"` //告警通知模版ID
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
	common.PageInfo
}

// EffectiveTime 生效时间
type EffectiveTime struct {
	Week      []string `json:"week"`
	StartTime int      `json:"startTime"`
	EndTime   int      `json:"endTime"`
}

type PrometheusConfig struct {
	ForDuration       int64   `json:"forDuration"`       //告警持续时间
	ComplexExpression string  `json:"complexExpression"` //复合条件
	Rules             []Rules `json:"rules"`             //告警规则
	IsUnionRule       int     `json:"isUnionRule"`       //逻辑判断条件（0:||,1:&&.2:复合条件：(1 AND 2) OR 3）
}

// Rules 告警条件规则
type Rules struct {
	TargetMapping    string  ` json:"targetMapping"`   // 告警指标映射
	TargetExpression string  `json:"targetExpression"` //告警指标表达式
	MetricName       string  ` json:"metricName"`      //指标名称
	Unit             string  ` json:"unit"`            // 告警指标单位
	Value            float64 ` json:"value"`           // 告警指标值
	Operator         string  ` json:"operator"`        // 告警操作符
	Severity         string  ` json:"severity"`        // 告警严重程度
	Description      string  ` json:"description"`     // 描述
}
