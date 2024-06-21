package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type DutyManagerRouter struct{}

func (DutyManagerRouter) InitDutyManagerRouter(Router *gin.RouterGroup) {
	dutyRouter := Router.Group("duty").Use(middleware.ParseTenant())
	dutyManagerApi := api.ApiGroupApp.SystemApiGroup.DutyManagerApi
	{
		dutyRouter.POST("", dutyManagerApi.Create)      //新增值班
		dutyRouter.GET("list", dutyManagerApi.List)     //查询值班列表
		dutyRouter.PUT("", dutyManagerApi.Update)       //查询值班列表
		dutyRouter.GET("delete", dutyManagerApi.Delete) //删除值班信息
	}
}
