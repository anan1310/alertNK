package repo

import (
	"alarm_collector/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type (
	NoticeRepo struct {
		entryRepo
	}
	interNoticeRepo interface {
		Get(r models.NoticeQuery) (models.AlertNotice, error)
		List(req models.NoticeQuery) ([]models.AlertNotice, int64, error)
		Create(r models.AlertNotice) error
		Update(r models.AlertNotice) error
		Delete(r models.NoticeQuery) error
	}
)

func newNoticeInterface(db *gorm.DB, g InterGormDBCli) interNoticeRepo {
	return &NoticeRepo{
		entryRepo{
			g:  g,
			db: db,
		}}
}
func (nr NoticeRepo) Get(r models.NoticeQuery) (models.AlertNotice, error) {
	var alertNoticeData models.AlertNotice
	db := nr.db.Model(&models.AlertNotice{}).Where("tenant_id = ? AND id = ?", r.TenantId, r.ID)
	if err := db.First(&alertNoticeData).Error; err != nil {
		return alertNoticeData, err
	}
	return alertNoticeData, nil
}

func (nr NoticeRepo) List(req models.NoticeQuery) ([]models.AlertNotice, int64, error) {
	var (
		alertNoticeObject []models.AlertNotice
		db                = nr.db.Model(&models.AlertNotice{})
		total             int64
	)
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db.Where("tenant_id = ?", req.TenantId)
	//查询总条数
	db.Count(&total)
	if err := db.Limit(limit).Offset(offset).Find(&alertNoticeObject).Error; err != nil {
		return nil, 0, err
	}
	return alertNoticeObject, total, nil
}

func (nr NoticeRepo) Create(r models.AlertNotice) error {
	var alertNotice models.AlertNotice
	if !errors.Is(nr.DB().Model(&models.AlertNotice{}).Where("tenant_id = ? and name = ?", r.TenantId, r.Name).First(&alertNotice).Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("告警通知对象 %s 已经存在", r.Name)
	} else {
		if err := nr.g.Create(models.AlertNotice{}, r); err != nil {
			return err
		}
	}
	return nil
}

func (nr NoticeRepo) Update(r models.AlertNotice) error {
	u := Updates{
		Table: models.AlertNotice{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.ID,
		},
		Omit:    "id",
		Updates: r,
	}
	if err := nr.g.Updates(u); err != nil {
		return err
	}
	return nil
}

func (nr NoticeRepo) Delete(r models.NoticeQuery) error {

	var ruleNum1, ruleNum2 int64
	db := nr.db.Model(&models.AlertRule{})
	db.Where("notice_id = ?", r.ID).Count(&ruleNum1)
	db.Where("notice_group LIKE ?", "%"+r.ID+"%").Count(&ruleNum2)
	if ruleNum1 != 0 || ruleNum2 != 0 {
		return fmt.Errorf("无法删除通知对象 %s, 因为已有告警规则绑定", r.ID)
	}

	d := Delete{
		Table: models.AlertNotice{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.ID,
		},
	}
	if err := nr.g.Delete(d); err != nil {
		return err
	}
	return nil
}
