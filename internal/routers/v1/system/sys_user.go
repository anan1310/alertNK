package system

import (
	"alarm_collector/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter 初始化用户
func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	useRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		useRouter.GET("list", userApi.List) //获取用户信息
	}

}
