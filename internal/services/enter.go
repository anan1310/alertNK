package services

import (
	"alarm_collector/internal/services/system"
	"alarm_collector/pkg/ctx"
)

var (
	UserService system.InterSysUserService
)

func NewServices(ctx *ctx.Context) {
	UserService = system.NewInterUserService(ctx)

}
