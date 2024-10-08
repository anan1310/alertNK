package system

import (
	"alarm_collector/api"
	"alarm_collector/middleware"
	"github.com/gin-gonic/gin"
)

type NoticeRouter struct{}

func (ur *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noticeRouter := Router.Group("notice").Use(middleware.ParseTenant())
	noticeApi := api.ApiGroupApp.SystemApiGroup.NoticeApi
	{
		noticeRouter.GET("list", noticeApi.List)         //获取通知对象
		noticeRouter.POST("", noticeApi.Create)          //新建通知对象
		noticeRouter.PUT("", noticeApi.Update)           //更新通知对象
		noticeRouter.GET("delete", noticeApi.Delete)     //删除通知对象
		noticeRouter.GET("getNoticeById", noticeApi.Get) //查询通知对象
	}

}
