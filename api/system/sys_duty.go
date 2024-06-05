package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type DutyManagerApi struct{}

func (DutyManagerApi) Create(ctx *gin.Context) {
	r := new(models.DutyManagement)
	response.BindJson(ctx, r)
	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Create(r)
	})
}

func (DutyManagerApi) List(ctx *gin.Context) {
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	tenantId := ctx.Query("tenantId")
	dutyManagementQuery := models.DutyManagementQuery{
		TenantId: tenantId,
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
	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Update(r)
	})
}

func (DutyManagerApi) Delete(ctx *gin.Context) {

	tenantId := ctx.Query("tenantId")
	dutyId := ctx.Query("dutyId")
	groupQuery := models.DutyManagementQuery{
		TenantId: tenantId,
		ID:       dutyId,
	}
	response.Service(ctx, func() (interface{}, interface{}) {
		return nil, services.DutyManagerService.Delete(&groupQuery)
	})
}
