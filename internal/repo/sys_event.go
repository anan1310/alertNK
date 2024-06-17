package repo

import (
	"alarm_collector/internal/models"
	"gorm.io/gorm"
)

type (
	EventRepo struct {
		entryRepo
	}

	interEventRepo interface {
		GetHistoryEvent(r models.AlertHisEventQuery) ([]models.AlertHisEvent, int64, error)
		CreateHistoryEvent(r models.AlertHisEvent) error
	}
)

func newEventInterface(db *gorm.DB, g InterGormDBCli) interEventRepo {
	return &EventRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (e EventRepo) GetHistoryEvent(r models.AlertHisEventQuery) ([]models.AlertHisEvent, int64, error) {
	var (
		historyEvents []models.AlertHisEvent
		db            = e.db.Model(&models.AlertHisEvent{})
		total         int64
	)

	db.Where("tenant_id = ?", r.TenantId)

	if r.DatasourceType != "" {
		db = db.Where("datasource_type = ?", r.DatasourceType)
	}

	if r.Severity != "" {
		db = db.Where("severity = ?", r.Severity)
	}

	if r.StartAt != 0 && r.EndAt != 0 {
		db = db.Where("first_trigger_time > ? and first_trigger_time < ?", r.StartAt, r.EndAt)
	}

	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Limit(limit).Offset(offset).Order("recover_time desc").Find(&historyEvents).Error; err != nil {
		return nil, 0, err
	}

	return historyEvents, total, nil

}

func (e EventRepo) CreateHistoryEvent(r models.AlertHisEvent) error {
	err := e.g.Create(models.AlertHisEvent{}, r)
	if err != nil {
		return err
	}

	return nil
}
