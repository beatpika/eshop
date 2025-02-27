package dal

import (
	"github.com/beatpika/eshop/app/token/biz/dal/redis"
	"github.com/beatpika/eshop/app/token/conf"
)

// Init init dal
func Init() error {
	// 初始化Redis
	redisCfg := conf.GetConf().Redis
	if err := redis.Init(redisCfg.Address, redisCfg.Username, redisCfg.Password, redisCfg.DB); err != nil {
		return err
	}
	return nil
}

// Close close dal
func Close() {
	redis.Close()
}
