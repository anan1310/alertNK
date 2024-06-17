package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AlertEventApi struct{}

func (AlertEventApi) ListCurrentEvent(ctx *gin.Context) {
	r := new(models.AlertCurEventQuery)

	response.BindJson(ctx, r)

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.EventService.ListCurrentEvent(r)
	})
}

func (AlertEventApi) ListHistoryEvent(ctx *gin.Context) {
	r := new(models.AlertHisEventQuery)
	response.BindJson(ctx, r)

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.EventService.ListHistoryEvent(r)
	})
}
