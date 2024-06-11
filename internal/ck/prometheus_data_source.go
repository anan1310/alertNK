package ck

import (
	"alarm_collector/internal/models"
	"gorm.io/gorm"
)

//TODO:prometheus 告警源数据-->告警数据获取

type (
	PrometheusDataSource struct {
		entryRepo
	}
	InterPrometheusSource interface {
		Get(models.PrometheusDataSourceQuery) (map[string]interface{}, error)
	}
)

func NewPrometheusDataSource(db *gorm.DB) InterPrometheusSource {
	return &PrometheusDataSource{
		entryRepo{
			db: db,
		},
	}
}

func (pds PrometheusDataSource) Get(r models.PrometheusDataSourceQuery) (map[string]interface{}, error) {
	var (
		db      = pds.db.Table("metric_" + r.MetricType)
		dataMap = make(map[string]interface{})
	)
	if err := db.Debug().Select(r.TargetMapping).Where("name = ? and pid = ? ", r.MetricName, r.Pid).Order("create_time desc").Limit(1).Scan(&dataMap).Error; err != nil {
		return nil, err
	}
	return dataMap, nil
}
