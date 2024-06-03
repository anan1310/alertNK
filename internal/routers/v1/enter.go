package v1

import "alarm_collector/internal/routers/v1/system"

type RouterGroup struct {
	SystemRouter system.RouterGroup //用户、角色、菜单
}

var RouterGroupApp = new(RouterGroup)
