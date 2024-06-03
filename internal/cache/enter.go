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
	}
)

func NewEntryCache() InterEntryCache {
	r := init_database.InitRedis()
	return &entryCache{
		redis: r,
	}
}

func (e entryCache) Redis() *redis.Client { return e.redis }
