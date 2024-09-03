package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type RuleRouter struct{}

func (RuleRouter) InitRuleRouter(Router *gin.RouterGroup) {
	rule := Router.Group("rule").Use(middleware.ParseTenant())
	ruleApi := api.ApiGroupApp.SystemApiGroup.RuleApi
	{
		rule.POST("", ruleApi.Create)            //添加告警规则
		rule.PUT("", ruleApi.Update)             //更新告警规则
		rule.GET("list", ruleApi.List)           //查询告警规则
		rule.GET("delete", ruleApi.Delete)       //删除告警规则
		rule.GET("getRuleById", ruleApi.Get)     //根据ID查询告警规则
		rule.GET("countRule", ruleApi.CountRule) //根据ID查询告警规则
	}
}
