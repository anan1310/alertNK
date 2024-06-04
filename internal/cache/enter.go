package cache

import (
	"alarm_collector/initialize/init_database"
	"github.com/go-redis/redis"
)

type (
	entryCache struct {
		redis *redis.Client
	}

	InterEntryCache interface {
		Redis() *redis.Client
		Silence() InterSilenceCache //静默规则
		Rule() InterRuleCache       //告警规则
		Event() InterEventCache
	}
)

func NewEntryCache() InterEntryCache {
	r := init_database.InitRedis()
	return &entryCache{
		redis: r,
	}
}

func (e entryCache) Redis() *redis.Client       { return e.redis }
func (e entryCache) Silence() InterSilenceCache { return newSilenceCacheInterface(e.redis) }
func (e entryCache) Rule() InterRuleCache       { return newRuleCacheInterface(e.redis) }
func (e entryCache) Event() InterEventCache     { return newEventCacheInterface(e.redis) }
