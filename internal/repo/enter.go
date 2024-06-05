package repo

import (
	"alarm_collector/initialize/init_database"
	"gorm.io/gorm"
)

type (
	entryRepo struct {
		g  interGormDBCli
		db *gorm.DB
	}

	InterEntryRepo interface {
		DB() *gorm.DB
		SysUser() interUserRepo
		RuleGroup() interRuleGroupRepo //告警组
		DutyManager() interDutyManagerRepo
		DutyCalendar() interDutyCalendar
	}
)

func NewMySQLRepoEntry() InterEntryRepo {
	db := init_database.Gorm("mysql")
	g := NewInterGormDBCli(db)
	return &entryRepo{
		g:  g,
		db: db,
	}
}

func (e *entryRepo) DB() *gorm.DB                      { return e.db }
func (e *entryRepo) SysUser() interUserRepo            { return newUserInterface(e.db, e.g) }
func (e *entryRepo) RuleGroup() interRuleGroupRepo     { return newRuleGroupInterface(e.db, e.g) }
func (e *entryRepo) DutyManager() interDutyManagerRepo { return newDutyManagerInterface(e.db, e.g) }
func (e *entryRepo) DutyCalendar() interDutyCalendar   { return newDutyCalendarInterface(e.db, e.g) }
