package models

// PrometheusDataSourceQuery prometheus请求数据源
type PrometheusDataSourceQuery struct {
	MetricType    string `json:"metricType"`
	MetricName    string `json:"metricName"`
	TargetMapping string `json:"targetMapping"` //所选指标
	Pid           string `json:"pid"`
}
