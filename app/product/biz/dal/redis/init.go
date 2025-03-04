package redis

import (
	"context"
	"time"

	"github.com/beatpika/eshop/app/product/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

// Init init redis client
func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		klog.Fatal("redis connect failed:", err)
	}
}

// GetRDB get redis client
func GetRDB() *redis.Client {
	return RDB
}
