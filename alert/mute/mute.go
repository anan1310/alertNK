package mute

import (
	"alarm_collector/global"
	"alarm_collector/internal/models"
	"alarm_collector/pkg/ctx"
	"time"
)

func IsMuted(ctx *ctx.Context, alert *models.AlertCurEvent, notice models.AlertNotice) bool {
	//判断是否开启通知 勾选之后才会通知

	if enable := IsNotificationEnabled(alert, notice); enable {
		return true
	}

	// 判断静默
	var as models.AlertSilences
	ctx.DB.DB().Model(models.AlertSilences{}).Where("fingerprint = ?", alert.Fingerprint).First(&as)

	_, ok := ctx.Redis.Silence().GetCache(models.AlertSilenceQuery{
		TenantId:    as.TenantId,
		Fingerprint: as.Fingerprint,
	})
	if ok {
		return true
	} else {
		ttl, _ := ctx.Redis.Redis().TTL(models.SilenceCachePrefix + alert.Fingerprint).Result()
		// 如果剩余生存时间小于0，表示键已过期
		if ttl < 0 {
			ctx.DB.DB().Model(models.AlertSilences{}).
				Where("tenant_id = ? AND fingerprint = ?", alert.TenantId, alert.Fingerprint).
				Delete(models.AlertSilences{})
		}
	}

	return InTheEffectiveTime(notice)
}

// InTheEffectiveTime 判断生效时间
func InTheEffectiveTime(notice models.AlertNotice) bool {
	if len(notice.UserNotices.Week) <= 0 {
		return false
	}

	var (
		p           bool
		currentTime = time.Now()
	)

	cwd := currentWeekday(currentTime)
	for _, wd := range notice.UserNotices.Week {
		if cwd != wd {
			continue
		}
		p = true
	}

	if !p {
		return true
	}

	cts := currentTimeSeconds(currentTime)
	if cts < notice.UserNotices.StartTime || cts > notice.UserNotices.EndTime {
		return true
	}

	return false
}

func currentWeekday(ct time.Time) string {
	// 获取当前时间
	currentDate := ct.Format("2006-01-02")

	// 解析日期字符串为时间对象
	date, err := time.Parse("2006-01-02", currentDate)
	if err != nil {
		global.Logger.Sugar().Error(err.Error())
		return ""
	}

	return date.Weekday().String()
}

func currentTimeSeconds(ct time.Time) int {
	cs := ct.Hour()*3600 + ct.Minute()*60
	return cs
}

func IsNotificationEnabled(alert *models.AlertCurEvent, notice models.AlertNotice) bool {
	switch alert.IsRecovered {
	case true:
		if !*notice.EnabledRecoverNotice {
			return true
		}
	case false:
		if !*notice.EnabledAlertNotice {
			return true
		}
	}
	return false
}
