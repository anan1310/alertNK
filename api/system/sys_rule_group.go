package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RuleGroupApi struct{}

func (RuleGroupApi) Create(ctx *gin.Context) {
	r := new(models.RuleGroups)
	response.BindJson(ctx, r)

	/*
		//之后设计将TenantID 存到请求头中 使用context进行一个管理
			tid, _ := ctx.Get("TenantID")
			r.TenantId = tid.(string)
	*/

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Create(r)
	})
}

func (RuleGroupApi) Update(ctx *gin.Context) {
	r := new(models.RuleGroups)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Update(r)
	})
}

func (RuleGroupApi) List(ctx *gin.Context) {
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	tenantId := ctx.Query("tenantId")
	groupQuery := models.RuleGroupQuery{
		TenantId: tenantId,
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.RuleGroupService.List(&groupQuery)
	})
}
func (RuleGroupApi) Delete(ctx *gin.Context) {

	tenantId := ctx.Query("tenantId")
	groupId := ctx.Query("groupId")
	groupQuery := models.RuleGroupQuery{
		TenantId: tenantId,
		ID:       groupId,
	}
	response.Service(ctx, func() (interface{}, interface{}) {
		err := services.RuleGroupService.Delete(&groupQuery)
		return nil, err
	})
}
