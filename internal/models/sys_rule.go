package models

import "alarm_collector/pkg/utils/common"

type LabelsMap map[string]string

type NoticeGroup []map[string]string

type AlertRule struct {
	TenantId             string   `json:"tenantId"`
	RuleId               string   `json:"ruleId" gorm:"ruleId"`                             //告警规则ID
	RuleGroupId          string   `json:"ruleGroupId"`                                      //所属告警组
	DatasourceType       string   `json:"datasourceType"`                                   //告警类型 Prometheus  Log  Apm
	MetricsParent        string   `json:"metricsParent"`                                    //告警大类
	MetricsChild         string   `json:"metricsChild"`                                     //告警小类
	DatasourceIdList     []string `json:"datasourceId" gorm:"datasourceId;serializer:json"` //数据源列表: 当前一个告警源对应一个数据源列表
	RuleName             string   `json:"ruleName"`                                         //规则名称
	EvalInterval         int64    `json:"evalInterval"`                                     //执行频率
	RepeatNoticeInterval int64    `json:"repeatNoticeInterval"`                             //重复通知间隔时间
	Description          string   `json:"description"`                                      //描述信息
	Severity             string   `json:"severity"`                                         //告警程度
	// Prometheus
	PrometheusConfig PrometheusConfig `json:"prometheusConfig" gorm:"prometheusConfig;serializer:json"` //prometheus相关配置

	NoticeId    string      `json:"noticeId"`
	NoticeGroup NoticeGroup `json:"noticeGroup" gorm:"noticeGroup;serializer:json"` //告警通知模版ID
	Enabled     *bool       `json:"enabled" gorm:"enabled"`                         //是否开启告警
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
	NoticeId         string   `json:"noticeId" form:"noticeId"`
	Enabled          string   `json:"enabled" form:"enabled"`
	common.PageInfo
}

type PrometheusConfig struct {
	ForDuration       int64             `json:"forDuration"`       //告警持续时间
	ComplexExpression string            `json:"complexExpression"` //复合条件
	AlertSource       map[string]string `json:"alertSource"`       //告警源
	Rules             []Rules           `json:"rules"`             //告警规则
	IsUnionRule       int               `json:"isUnionRule"`       //逻辑判断条件（0:||,1:&&.2:复合条件：(1 AND 2) OR 3）
}

// Rules 告警条件规则
type Rules struct {
	TargetKey        string  `json:"targetKey"`        //唯一标识
	TargetMapping    string  ` json:"targetMapping"`   // 告警指标映射 “memory_available_bytes,memory_total_bytes”
	TargetExpression string  `json:"targetExpression"` // 告警指标表达式  "(1- (@[memory_available_bytes]@) / @[memory_total_bytes]@) * 100"
	MetricName       string  ` json:"metricName"`      // 指标名称
	FromUnit         string  `json:"fromUnit"`         //告警原单位
	ToUnit           string  `json:"toUnit"`           //告警目标单位
	Precision        string  `json:"precision"`        //保留精度（默认2位）
	Value            float64 ` json:"value"`           // 告警阈值
	ValueMax         float64 `json:"valueMax"`         //最大值
	ValueMin         float64 `json:"valueMin"`         //最小值
	Operator         string  ` json:"operator"`        // 告警操作符
	Severity         string  ` json:"severity"`        // 告警严重程度
	Description      string  ` json:"description"`     // 描述 "内存使用率"
}

// AlertScreen 大屏告警相关
type AlertScreen struct {
	//当前告警数
	RecentCount int `json:"recentRuleCount"`
	//历史告警数
	HistoryCount int64 `json:"historyCount"`
	//告警规则数
	RuleCount int64 `json:"ruleCount"`
}
