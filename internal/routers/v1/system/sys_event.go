package system

import (
	"alarm_collector/api"
	"github.com/gin-gonic/gin"
)

type AlertEventRouter struct{}

func (AlertEventRouter) InitAlertEventRouter(Router *gin.RouterGroup) {
	eventRouter := Router.Group("event")
	alertEvenApi := api.ApiGroupApp.SystemApiGroup.AlertEventApi
	{
		eventRouter.POST("curEvent", alertEvenApi.ListCurrentEvent)     //获取当前告警信息
		eventRouter.POST("historyEvent", alertEvenApi.ListHistoryEvent) //获取历史告警信息
	}
}
