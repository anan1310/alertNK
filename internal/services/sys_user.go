package services

import (
	"alarm_collector/pkg/ctx"
)

type sysUserService struct {
	ctx *ctx.Context
}

type interSysUserService interface {
	List() (interface{}, interface{})
}

func newInterUserService(ctx *ctx.Context) interSysUserService {
	return &sysUserService{
		ctx: ctx,
	}
}

func (us sysUserService) List() (interface{}, interface{}) {
	data, err := us.ctx.DB.SysUser().List()
	if err != nil {
		return nil, err
	}
	return data, nil
}
