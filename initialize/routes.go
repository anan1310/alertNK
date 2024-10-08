package initialize

import (
	"alarm_collector/global"
	v1 "alarm_collector/internal/routers/v1"
	"alarm_collector/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func RunServer() {
	gin.SetMode(global.Config.Server.RunMode)
	//Debug:调试版本，包含调试信息，容量比Release大很多
	//Release:发布版本，不对源代码进行调试
	r := routersInit() //初始化路由
	readTimeout := global.Config.Server.ReadTimeout * time.Second
	writeTimeout := global.Config.Server.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", global.Config.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}

func routersInit() *gin.Engine {

	r := gin.New()
	/*
		// 初始化主库
		masterDB := init_database.Gorm("mysql")
		//初始化Redis
		rCache := cache.NewEntryCache()
		//初始化ClickHouse
		ckRepo := ck.NewClickHouseRepoEntry()
	*/
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": "true",
		})
	})
	// 解决跨域的问题
	r.Use(
		// 启用CORS中间件
		middleware.Cors(),
		// 拦截panic异常
		middleware.Recovery(false),
		// 自定义请求日志格式
		gin.LoggerWithFormatter(middleware.RequestLoggerFormatter),
		//多数据源处理中间件
		//middleware.ParseTenant(),
		//middleware.MilTenant(masterDB, rCache, ckRepo),
	)

	PrivateGroup := r.Group("system")
	{
		v1.RouterGroupApp.SystemRouter.InitUserRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitRuleGroupRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitRuleRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitDutyManagerRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitDutyCalendarRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitNoticeRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitSilencesRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitAlertEventRouter(PrivateGroup)
		v1.RouterGroupApp.SystemRouter.InitDashBoardInfoRouter(PrivateGroup)
	}

	return r

}
