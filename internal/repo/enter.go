package repo

import (
	"alarm_collector/initialize/init_database"
	"gorm.io/gorm"
)

type (
	entryRepo struct {
		g  InterGormDBCli
		db *gorm.DB
	}

	InterEntryRepo interface {
		DB() *gorm.DB
		SysUser() InterUserRepo
		RuleGroup() InterRuleGroupRepo //告警组
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

func (e *entryRepo) DB() *gorm.DB                  { return e.db }
func (e *entryRepo) SysUser() InterUserRepo        { return newUserInterface(e.db, e.g) }
func (e *entryRepo) RuleGroup() InterRuleGroupRepo { return newRuleGroupInterface(e.db, e.g) }
