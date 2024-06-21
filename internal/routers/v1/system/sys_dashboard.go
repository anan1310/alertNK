package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type DashBoardInfoRouter struct{}

func (DashBoardInfoRouter) InitDashBoardInfoRouter(Router *gin.RouterGroup) {
	dashBoardRouter := Router.Group("dashBoard").Use(middleware.ParseTenant())
	dashBoardInfoApi := api.ApiGroupApp.SystemApiGroup.DashBoardInfoApi
	{
		dashBoardRouter.GET("dashBoardInfo", dashBoardInfoApi.GetDashBoardInfo) //告警面板
	}
}
