package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RuleApi struct{}

func (RuleApi) Create(ctx *gin.Context) {
	r := new(models.AlertRule)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Create(r)
	})
}

func (RuleApi) Update(ctx *gin.Context) {
	r := new(models.AlertRule)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Update(r)
	})
}

func (RuleApi) List(ctx *gin.Context) {
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	tenantId := ctx.Query("tenantId")
	ruleGroupId := ctx.Query("ruleGroupId")
	ruleQuery := models.AlertRuleQuery{
		TenantId:    tenantId,
		RuleGroupId: ruleGroupId,
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.RuleService.ListRule(&ruleQuery)
	})
}

func (RuleApi) Delete(ctx *gin.Context) {
	tenantId := ctx.Query("tenantId")
	ruleId := ctx.Query("ruleId")
	ruleGroupId := ctx.Query("ruleGroupId")
	groupQuery := models.AlertRuleQuery{
		TenantId:    tenantId,
		RuleId:      ruleId,
		RuleGroupId: ruleGroupId,
	}
	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Delete(&groupQuery)
	})
}
