package repo

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/cmd"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	DutyManagerRepo struct {
		entryRepo
	}
	interDutyManagerRepo interface {
		Create(r models.DutyManagement) error
		Update(r models.DutyManagement) error
		List(req models.DutyManagementQuery) ([]models.DutyManagement, int64, error)
		Delete(req models.DutyManagementQuery) error
	}
)

func newDutyManagerInterface(db *gorm.DB, g InterGormDBCli) interDutyManagerRepo {
	return &DutyManagerRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (d DutyManagerRepo) Create(r models.DutyManagement) error {
	nr := r
	nr.ID = "dt-" + cmd.RandId()
	nr.CreateAt = time.Now().Unix()
	err := d.g.Create(&models.DutyManagement{}, nr)
	if err != nil {
		return err
	}
	return nil
}

func (d DutyManagerRepo) List(req models.DutyManagementQuery) ([]models.DutyManagement, int64, error) {
	var (
		data  []models.DutyManagement
		db    = d.db.Model(&models.DutyManagement{})
		total int64
	)
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db.Model(&models.DutyManagement{}).Where("tenant_id = ?", req.TenantId)
	//查询总条数
	db.Count(&total)
	if err := db.Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (d DutyManagerRepo) Update(r models.DutyManagement) error {
	u := Updates{
		Table: models.DutyManagement{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.ID,
		},
		Omit:    "id",
		Updates: r,
	}
	if err := d.g.Updates(u); err != nil {
		return err
	}
	return nil
}

func (d DutyManagerRepo) Delete(r models.DutyManagementQuery) error {
	var noticeNum int64
	db := d.db.Model(&models.AlertNotice{})
	db.Where("tenant_id = ? AND duty_id = ?", r.TenantId, r.ID).Count(&noticeNum)
	if noticeNum != 0 {
		return fmt.Errorf("无法删除值班表 %s, 因为已有通知对象绑定", r.ID)
	}
	//删除值班表
	delDuty := Delete{
		Table: models.DutyManagement{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"id = ?":        r.ID,
		},
	}
	err := d.g.Delete(delDuty)
	if err != nil {
		return err
	}
	//删除值班表每条信息
	delCalendar := Delete{
		Table: models.DutySchedule{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"duty_id = ?":   r.ID,
		},
	}
	if err := d.g.Delete(delCalendar); err != nil {
		return err
	}
	return nil
}
