package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type SilencesApi struct{}

func (SilencesApi) Create(ctx *gin.Context) {
	r := new(models.AlertSilences)
	response.BindJson(ctx, r)

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Create(r)
	})
}

func (SilencesApi) Update(ctx *gin.Context) {
	r := new(models.AlertSilences)
	response.BindJson(ctx, r)

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Update(r)
	})
}

func (SilencesApi) Delete(ctx *gin.Context) {
	r := new(models.AlertSilenceQuery)
	response.BindQuery(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Delete(r)
	})
}

func (SilencesApi) List(ctx *gin.Context) {
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	silenceQuery := &models.AlertSilenceQuery{
		TenantId: tid.(string),
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.SilenceService.ListSilence(silenceQuery)
	})
}
