package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AlertEventApi struct{}

func (AlertEventApi) ListCurrentEvent(ctx *gin.Context) {
	r := new(models.AlertCurEventQuery)

	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.EventService.ListCurrentEvent(r)
	})
}

func (AlertEventApi) ListHistoryEvent(ctx *gin.Context) {
	r := new(models.AlertHisEventQuery)
	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.EventService.ListHistoryEvent(r)
	})
}
