package init_database

import (
	"alarm_collector/global"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var Redis *redis.Client

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Pass,
		DB:       0, // 使用默认的数据库
	})

	// 尝试连接到 Redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		global.Logger.Error("redis Connection Failed", zap.Error(err))
		return nil
	}

	Redis = client

	return client

}
