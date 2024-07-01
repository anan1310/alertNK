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
		List(userIds []int) ([]system.SysUser, error)
	}
)

func newUserInterface(db *gorm.DB, g InterGormDBCli) interUserRepo {
	return &UserRepo{
		entryRepo{
			g:  g,
			db: db,
		},
	}
}

func (u UserRepo) List(userIds []int) ([]system.SysUser, error) {
	var userList []system.SysUser
	err := u.db.Where("del_flag = ? and user_id in (?)", "0", userIds).Find(&userList).Error
	return userList, err
}
