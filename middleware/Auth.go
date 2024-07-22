package middleware

import (
	jwtUtils "alarm_collector/pkg/utils/jwt"
	"alarm_collector/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// Bearer Token, 获取 Token 值
		tokenStr = tokenStr[len(jwtUtils.TokenType)+1:]

		// 校验 Token
		code, ok := jwtUtils.IsTokenValid(tokenStr)
		if !ok {
			if code == 401 {
				response.TokenFail(c)
				c.Abort()
				return
			}
		}
	}
}
