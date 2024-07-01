package services

import (
	"alarm_collector/pkg/ctx"
)

type sysUserService struct {
	ctx *ctx.Context
}

type interSysUserService interface {
	List(userIds []int) (interface{}, interface{})
}

func newInterUserService(ctx *ctx.Context) interSysUserService {
	return &sysUserService{
		ctx: ctx,
	}
}

func (us sysUserService) List(userIds []int) (interface{}, interface{}) {
	data, err := us.ctx.DB.SysUser().List(userIds)
	if err != nil {
		return nil, err
	}
	return data, nil
}
