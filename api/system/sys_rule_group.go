package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type RuleGroupApi struct{}

func (RuleGroupApi) Create(ctx *gin.Context) {
	r := new(models.RuleGroups)
	response.BindJson(ctx, r)

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Create(r)
	})
}

func (RuleGroupApi) Update(ctx *gin.Context) {
	r := new(models.RuleGroups)
	response.BindJson(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Update(r)
	})
}

func (RuleGroupApi) List(ctx *gin.Context) {
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)

	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	groupQuery := models.RuleGroupQuery{
		TenantId: tid.(string),
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
	r := new(models.RuleGroupQuery)
	response.BindQuery(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		err := services.RuleGroupService.Delete(r)
		return nil, err
	})
}
