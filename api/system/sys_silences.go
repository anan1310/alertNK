package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type SilencesApi struct{}

func (SilencesApi) Create(ctx *gin.Context) {
	r := new(models.AlertSilences)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Create(r)
	})
}

func (SilencesApi) Update(ctx *gin.Context) {
	r := new(models.AlertSilences)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Update(r)
	})
}

func (SilencesApi) Delete(ctx *gin.Context) {
	tenantId := ctx.Query("tenantId")
	silenceId := ctx.Query("silenceId")
	silenceQuery := models.AlertSilenceQuery{
		TenantId: tenantId,
		ID:       silenceId,
	}

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Delete(&silenceQuery)
	})
}

func (SilencesApi) List(ctx *gin.Context) {
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	tenantId := ctx.Query("tenantId")
	silenceQuery := &models.AlertSilenceQuery{
		TenantId: tenantId,
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.SilenceService.ListSilence(silenceQuery)
	})
}
