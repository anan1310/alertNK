package models

// PrometheusDataSourceQuery prometheus请求数据源
type PrometheusDataSourceQuery struct {
	MetricType    string `json:"metricType"`
	MetricName    string `json:"metricName"`
	MetricHost    string `json:"metricHost"`
	TargetMapping string `json:"targetMapping"` //所选指标
	TenantId      string `json:"tenantId"`      //租户ID
}
