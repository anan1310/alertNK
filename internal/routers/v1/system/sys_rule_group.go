package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type RuleGroupRouter struct{}

func (RuleGroupRouter) InitRuleGroupRouter(Router *gin.RouterGroup) {
	ruleGroup := Router.Group("group").Use(middleware.ParseTenant())
	ruleGroupApi := api.ApiGroupApp.SystemApiGroup.RuleGroupApi
	{
		ruleGroup.POST("", ruleGroupApi.Create)      //添加告警组规则
		ruleGroup.PUT("", ruleGroupApi.Update)       //更新告警组规则
		ruleGroup.GET("list", ruleGroupApi.List)     //查询告警组规则
		ruleGroup.GET("delete", ruleGroupApi.Delete) //删除告警组规则
	}
}
