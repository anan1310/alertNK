package system

import (
	"alarm_collector/internal/services"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserApi struct{}

func (uc UserApi) List(ctx *gin.Context) {

	us := ctx.Query("userIds")
	var userIds []int
	for _, id := range strings.Split(us, ",") {
		userIds = append(userIds, common.ToInt(id))
	}

	response.Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.List(userIds)
	})
}
