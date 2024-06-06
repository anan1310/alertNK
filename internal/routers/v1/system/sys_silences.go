package system

import (
	"alarm_collector/api"
	"github.com/gin-gonic/gin"
)

type SilencesRouter struct{}

func (sr *SilencesRouter) InitSilencesRouter(Router *gin.RouterGroup) {
	silencesRouter := Router.Group("silences")
	silencesApi := api.ApiGroupApp.SystemApiGroup.SilencesApi
	{
		silencesRouter.GET("list", silencesApi.List)     //获取静默规则
		silencesRouter.POST("", silencesApi.Create)      //新建静默规则
		silencesRouter.PUT("", silencesApi.Update)       //更新静默规则
		silencesRouter.GET("delete", silencesApi.Delete) //删除静默规则
	}

}
