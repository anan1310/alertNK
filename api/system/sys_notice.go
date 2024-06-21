package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type NoticeApi struct{}

func (NoticeApi) List(ctx *gin.Context) {
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	NoticeQuery := &models.NoticeQuery{
		TenantId: tid.(string),
		PageInfo: common.PageInfo{
			Page:     page,
			PageSize: pageSize,
		},
	}

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.NoticeService.List(NoticeQuery)
	})
}

func (NoticeApi) Create(ctx *gin.Context) {
	r := new(models.AlertNotice)
	response.BindJson(ctx, r)

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Create(r)
	})
}

func (NoticeApi) Update(ctx *gin.Context) {
	r := new(models.AlertNotice)
	response.BindJson(ctx, r)
	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Update(r)
	})
}

func (NoticeApi) Delete(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	response.BindQuery(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Delete(r)
	})
}

func (NoticeApi) Get(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	response.BindQuery(ctx, r)

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Get(r)
	})

}
