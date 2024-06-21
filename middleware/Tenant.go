package middleware

import (
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

const TenantIDHeaderKey = "TenantID"

func ParseTenant() gin.HandlerFunc {
	// 从HTTP头部获取TenantID并存储到上下文中，可以提高代码的可维护性、可重用性、安全性和性能，同时也使得错误处理和业务逻辑的实现更加高效和灵活。
	return func(context *gin.Context) {
		tid := context.Request.Header.Get(TenantIDHeaderKey)
		if tid == "" {
			response.Fail(context, "您还没有选择项目", "failed")
			context.Abort()
			return
		}

		context.Set(TenantIDHeaderKey, tid)
		context.Next()
	}
}
