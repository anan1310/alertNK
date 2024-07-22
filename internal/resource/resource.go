package resource

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"runtime"
	"time"
)

func InitResource(ctx *ctx.Context) {
	ticker := time.Tick(time.Second * 60)
	go func() {
		for range ticker {
			curAt := time.Now()
			goNum := runtime.NumGoroutine()
			cleanupOldData(curAt)
			ctx.DB.DB().Model(&models.ServiceResource{}).Create(models.ServiceResource{
				ID:    uint(curAt.Unix()),
				Time:  curAt.Format(global.Layout),
				Value: goNum,
				Label: "协程数",
			})
		}
	}()

}

func cleanupOldData(curAt time.Time) {
	c := ctx.DO()
	cutoffTime := curAt.Add(-6 * time.Hour)
	c.DB.DB().Where("time < ?", cutoffTime).Delete(&models.ServiceResource{})
}
