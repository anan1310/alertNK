package system

import (
	"alarm_collector/internal/models"
	"alarm_collector/internal/services"
	"alarm_collector/middleware"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type DutyCalendarApi struct{}

func (DutyCalendarApi) Create(ctx *gin.Context) {
	r := new(models.DutyScheduleCreate)
	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.DutyCalendarService.CreateAndUpdate(r)
	})
}

func (DutyCalendarApi) Update(ctx *gin.Context) {
	r := new(models.DutySchedule)
	err := response.BindJson(ctx, r)
	if err != nil {
		return
	}

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.DutyCalendarService.Update(r)
	})
}

func (DutyCalendarApi) List(ctx *gin.Context) {
	r := new(models.DutyScheduleQuery)
	err := response.BindQuery(ctx, r)
	if err != nil {
		return
	}

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.DutyCalendarService.List(r)
	})
}
func (DutyCalendarApi) GetDutyUserInfo(ctx *gin.Context) {
	r := new(models.DutyScheduleQuery)
	err := response.BindQuery(ctx, r)
	if err != nil {
		return
	}

	//存到请求头中 使用context进行一个管理
	tid, _ := ctx.Get(middleware.TenantIDHeaderKey)
	r.TenantId = tid.(string)

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.DutyCalendarService.GetDutyUserInfo(r), nil
	})
}
