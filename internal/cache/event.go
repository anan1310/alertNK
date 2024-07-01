package cache

import (
	"alarm_collector/global"
	"alarm_collector/initialize/init_database"
	"alarm_collector/internal/models"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

type (
	eventCache struct {
		rc *redis.Client
	}

	InterEventCache interface {
		SetCache(cacheType string, event models.AlertCurEvent, expiration time.Duration)
		DelCache(key string)
		GetCache(key string) models.AlertCurEvent
		GetFirstTime(key string) int64
		GetLastEvalTime(key string) int64
		GetLastSendTime(key string) int64
	}
)

func newEventCacheInterface(r *redis.Client) InterEventCache {
	return &eventCache{
		r,
	}
}

func (ec eventCache) SetCache(cacheType string, event models.AlertCurEvent, expiration time.Duration) {
	alertJson, _ := json.Marshal(event)
	switch cacheType {
	case "Firing":
		init_database.Redis.Set(event.GetFiringAlertCacheKey(), string(alertJson), expiration)
	case "Pending":
		init_database.Redis.Set(event.GetPendingAlertCacheKey(), string(alertJson), expiration)
	}

}

func (ec eventCache) DelCache(key string) {
	// 使用Scan命令获取所有匹配指定模式的键
	iter := init_database.Redis.Scan(0, key, 0).Iterator()
	keysToDelete := make([]string, 0)

	// 遍历匹配的键
	for iter.Next() {
		key := iter.Val()
		keysToDelete = append(keysToDelete, key)
	}

	if err := iter.Err(); err != nil {
		global.Logger.Sugar().Error("redis find keys error")
	}

	// 批量删除键
	if len(keysToDelete) > 0 {
		err := init_database.Redis.Del(keysToDelete...).Err()
		if err != nil {
			global.Logger.Sugar().Error("redis del err:", err)
		}
		global.Logger.Sugar().Infof("移除告警消息 -> %s\n", keysToDelete)
	}
}

func (ec eventCache) GetCache(key string) models.AlertCurEvent {

	var alert models.AlertCurEvent

	d, err := ec.rc.Get(key).Result()
	_ = json.Unmarshal([]byte(d), &alert)
	if err != nil {
		return alert
	}

	return alert

}

func (ec eventCache) GetFirstTime(key string) int64 {

	ft := ec.GetCache(key).FirstTriggerTime
	if ft == 0 {
		return time.Now().Unix()
	}
	return ft

}

func (ec eventCache) GetLastEvalTime(key string) int64 {

	curTime := time.Now().Unix()
	let := ec.GetCache(key).LastEvalTime
	if let == 0 || let < curTime {
		return curTime
	}

	return let

}

func (ec eventCache) GetLastSendTime(key string) int64 {

	return ec.GetCache(key).LastSendTime

}
