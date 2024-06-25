package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type DutyManagerApi struct{}

func (DutyManagerApi) Create(ctx *gin.Context) {
	r := new(models.DutyManagement)
	response.BindJson(ctx, r)
	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Create(r)
	})
}

func (DutyManagerApi) List(ctx *gin.Context) {
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)

	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	name := ctx.Query("name")
	dutyManagementQuery := models.DutyManagementQuery{
		TenantId: tid.(string),
		Name:     name,
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}
	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.DutyManagerService.List(&dutyManagementQuery)
	})
}

func (DutyManagerApi) Update(ctx *gin.Context) {
	r := new(models.DutyManagement)
	response.BindJson(ctx, r)

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)
	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Update(r)
	})
}

func (DutyManagerApi) Delete(ctx *gin.Context) {

	r := new(models.DutyManagementQuery)
	response.BindQuery(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Delete(r)
	})
}
