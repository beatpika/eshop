package conf

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// GetRedisClient returns the Redis client instance
func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		conf := GetConf()
		redisClient = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Address,
			Username: conf.Redis.Username,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		})
	})
	return redisClient
}
