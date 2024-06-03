package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
	})

}
