package initialize

import (
	"alarm_collector/alert"
	"alarm_collector/core"
	"alarm_collector/global"
	"alarm_collector/internal/cache"
	"alarm_collector/internal/ck"
	"alarm_collector/internal/repo"
	"alarm_collector/internal/resource"
	"alarm_collector/internal/services"
	"alarm_collector/pkg/ctx"
	"context"
)

func InitBasic() {

	// 初始化配置
	global.Viper = core.Viper()
	//初始化日志配置
	global.Logger = core.Zap()
	//初始化MySQL
	dbRepo := repo.NewMySQLRepoEntry()
	//初始化Redis
	rCache := cache.NewEntryCache()
	//初始化ClickHouse
	ckRepo := ck.NewClickHouseRepoEntry()

	newContext := ctx.NewContext(context.Background(), dbRepo, rCache, ckRepo)
	services.NewServices(newContext)
	// 启用告警评估协程
	alert.Initialize(newContext)
	//  初始化监控分析数据 (统计协程数)
	resource.InitResource(newContext)

}
