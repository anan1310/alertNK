package ck

import (
	"alarm_collector/initialize/init_database"
	"gorm.io/gorm"
)

type (
	entryRepo struct {
		db *gorm.DB
	}

	InterEntryRepo interface {
		DB() *gorm.DB
		PrometheusDataSource() InterPrometheusSource
	}
)

func NewClickHouseRepoEntry() InterEntryRepo {
	db := init_database.Gorm("clickhouse")
	return &entryRepo{
		db: db,
	}
}

func (e *entryRepo) DB() *gorm.DB { return e.db }
func (e *entryRepo) PrometheusDataSource() InterPrometheusSource {
	return NewPrometheusDataSource(e.db)
}
