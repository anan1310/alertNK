package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type NoticeApi struct{}

func (NoticeApi) List(ctx *gin.Context) {

	page := common.ToInt(ctx.Query("page"))
	pageSize := common.ToInt(ctx.Query("pageSize"))
	tenantId := ctx.Query("tenantId")
	NoticeQuery := &models.NoticeQuery{
		TenantId: tenantId,
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

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Create(r)
	})
}

func (NoticeApi) Update(ctx *gin.Context) {
	r := new(models.AlertNotice)
	response.BindJson(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Update(r)
	})
}

func (NoticeApi) Delete(ctx *gin.Context) {
	tenantId := ctx.Query("tenantId")
	noticeIdId := ctx.Query("noticeId")
	noticeQuery := &models.NoticeQuery{
		TenantId: tenantId,
		ID:       noticeIdId,
	}

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Delete(noticeQuery)
	})
}

func (NoticeApi) Get(ctx *gin.Context) {
	r := new(models.NoticeQuery)
	response.BindQuery(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.NoticeService.Get(r)
	})

}
