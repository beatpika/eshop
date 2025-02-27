package redis

import (
	"github.com/redis/go-redis/v9"
)

var (
	// Client Redis客户端实例
	Client *redis.Client
	// TokenManager Token存储管理器实例
	TokenManager *TokenStore
)

// Init 初始化Redis连接
func Init(address string, username string, password string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	TokenManager = NewTokenStore(Client)
	return nil
}

// Close 关闭Redis连接
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}
