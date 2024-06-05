package system

import (
	"alarm_collector/api"
	"github.com/gin-gonic/gin"
)

type DutyCalendarRouter struct{}

func (DutyManagerRouter) InitDutyCalendarRouter(Router *gin.RouterGroup) {
	calendarRouter := Router.Group("calendar")
	dutyCalendarApi := api.ApiGroupApp.SystemApiGroup.DutyCalendarApi
	{
		calendarRouter.POST("", dutyCalendarApi.Create)  //发布日程
		calendarRouter.GET("list", dutyCalendarApi.List) //查询日程
		calendarRouter.PUT("", dutyCalendarApi.Update)   //更新日程
	}
}
