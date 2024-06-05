package repo

import (
	"alarm_collector/internal/models/system"
	"gorm.io/gorm"
)

type (
	UserRepo struct {
		entryRepo
	}
	interUserRepo interface {
		List() ([]system.SysUser, error)
	}
)

func newUserInterface(db *gorm.DB, g interGormDBCli) interUserRepo {
	return &UserRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (u UserRepo) List() ([]system.SysUser, error) {
	var userList []system.SysUser
	err := u.db.Where("del_flag = ?", "0").Find(&userList).Error
	return userList, err
}
