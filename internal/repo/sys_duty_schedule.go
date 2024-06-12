package repo

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/internal/models/system"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	DutyCalendarRepo struct {
		entryRepo
	}

	interDutyCalendar interface {
		GetDutyUserInfo(dutyId, time string) system.SysUser
		GetCalendarInfo(dutyId, time string) models.DutySchedule
		Create(r models.DutySchedule) error
		Update(r models.DutySchedule) error
		List(r models.DutyScheduleQuery) ([]models.DutySchedule, error)
	}
)

func newDutyCalendarInterface(db *gorm.DB, g InterGormDBCli) interDutyCalendar {
	return &DutyCalendarRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (dc DutyCalendarRepo) GetCalendarInfo(dutyId, time string) models.DutySchedule {
	var dutySchedule models.DutySchedule

	if err := dc.db.Model(models.DutySchedule{}).
		Where("duty_id = ? AND time = ?", dutyId, time).
		First(&dutySchedule).Error; err != nil {
		global.Logger.Sugar().Error("get calendar info error", err.Error())
	}

	return dutySchedule
}

// GetDutyUserInfo 获取值班用户信息
func (dc DutyCalendarRepo) GetDutyUserInfo(dutyId, time string) system.SysUser {
	var user system.SysUser

	schedule := dc.GetCalendarInfo(dutyId, time)

	dc.db.Model(system.SysUser{}).
		Where("user_id = ?", schedule.DutyUser.UserId).
		First(&user)

	return user
}

func (dc DutyCalendarRepo) Create(r models.DutySchedule) error {
	if err := dc.g.Create(models.DutySchedule{}, r); err != nil {
		return err
	}
	return nil
}

func (dc DutyCalendarRepo) Update(r models.DutySchedule) error {
	u := Updates{
		Table: models.DutySchedule{},
		Where: map[string]interface{}{
			"tenant_id = ?": r.TenantId,
			"duty_id = ?":   r.DutyId,
			"time = ?":      r.Time,
		},
		Updates: r,
	}

	if err := dc.g.Updates(u); err != nil {
		return err
	}
	return nil
}

func (dc DutyCalendarRepo) List(r models.DutyScheduleQuery) ([]models.DutySchedule, error) {
	var dutyScheduleList []models.DutySchedule
	db := dc.db.Model(&models.DutySchedule{})

	if r.Time != "" {
		db.Where("tenant_id = ? AND duty_id = ? AND time Like ?", r.TenantId, r.DutyId, r.Time+"%").Find(&dutyScheduleList)
		return dutyScheduleList, nil
	}

	yearMonth := fmt.Sprintf("%d-%d-", time.Now().Year(), time.Now().Month())
	db.Where("tenant_id = ? AND duty_id = ? AND time LIKE ?", r.TenantId, r.DutyId, yearMonth+"%")
	err := db.Find(&dutyScheduleList).Error
	if err != nil {
		return dutyScheduleList, err
	}

	return dutyScheduleList, nil
}
