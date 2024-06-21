package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type DutyCalendarRouter struct{}

func (DutyManagerRouter) InitDutyCalendarRouter(Router *gin.RouterGroup) {
	calendarRouter := Router.Group("calendar").Use(middleware.ParseTenant())
	dutyCalendarApi := api.ApiGroupApp.SystemApiGroup.DutyCalendarApi
	{
		calendarRouter.GET("list", dutyCalendarApi.List)                       //查询日程
		calendarRouter.POST("", dutyCalendarApi.Create)                        //发布日程
		calendarRouter.PUT("", dutyCalendarApi.Update)                         //更新日程
		calendarRouter.GET("getDutyUserInfo", dutyCalendarApi.GetDutyUserInfo) //获取日程信息
	}
}
