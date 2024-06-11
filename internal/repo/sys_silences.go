package repo

import (
	"alarm_collector/internal/models"
	"gorm.io/gorm"
)

// TODO 告警静默
type (
	SilenceRepo struct {
		entryRepo
	}

	interSilenceRepo interface {
		List(r models.AlertSilenceQuery) ([]models.AlertSilences, int64, error)
		Create(r models.AlertSilences) error
		Update(r models.AlertSilences) error
		Delete(r models.AlertSilenceQuery) error
	}
)

func newSilenceInterface(db *gorm.DB, g InterGormDBCli) interSilenceRepo {
	return &SilenceRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (sr SilenceRepo) List(r models.AlertSilenceQuery) ([]models.AlertSilences, int64, error) {
	var (
		silenceList []models.AlertSilences
		total       int64
		db          = sr.db.Model(models.AlertSilences{})
	)
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db.Where("tenant_id = ?", r.TenantId)
	db.Count(&total)

	err := db.Limit(limit).Offset(offset).Find(&silenceList).Error
	if err != nil {
		return silenceList, 0, err
	}

	return silenceList, total, nil
}

func (sr SilenceRepo) Create(r models.AlertSilences) error {
	err := sr.g.Create(models.AlertSilences{}, r)
	if err != nil {
		return err
	}

	return nil
}

func (sr SilenceRepo) Update(r models.AlertSilences) error {
	u := Updates{
		Table: models.AlertSilences{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.Id,
		},
		Omit:    "id",
		Updates: r,
	}

	err := sr.g.Updates(u)
	if err != nil {
		return err
	}

	return nil
}

func (sr SilenceRepo) Delete(r models.AlertSilenceQuery) error {
	var silence models.AlertSilences
	db := sr.db.Where("tenant_id = ? AND id = ?", r.TenantId, r.ID)
	err := db.Find(&silence).Error
	if err != nil {
		return err
	}

	del := Delete{
		Table: models.AlertSilences{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.ID,
		},
	}
	err = sr.g.Delete(del)
	if err != nil {
		return err
	}

	return nil
}
