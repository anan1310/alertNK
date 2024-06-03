package system

import (
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (uc UserApi) List(c *gin.Context) {
	response.HandleResponse(c, func() (interface{}, interface{}) {
		return services.UserService.List()
	})
}
