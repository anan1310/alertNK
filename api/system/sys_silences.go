package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
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
	r := new(models.AlertSilenceQuery)
	response.BindQuery(ctx, r)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Delete(r)
	})
}

func (SilencesApi) List(ctx *gin.Context) {
	r := new(models.AlertSilenceQuery)
	response.BindQuery(ctx, r)

	response.ServiceTotal(ctx, func() (interface{}, interface{}, interface{}) {
		return services.SilenceService.List(r)
	})
}
