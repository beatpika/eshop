package service

import (
	"sync"

	"github.com/beatpika/eshop/app/token/biz/utils"
	"github.com/beatpika/eshop/app/token/conf"
)

var (
	// jwtUtil 全局JWT工具实例
	jwtUtil  *utils.JWTUtil
	initOnce sync.Once
)

// initJWTUtil 延迟初始化JWT工具
func initJWTUtil() {
	initOnce.Do(func() {
		jwtUtil = utils.NewJWTUtil(conf.GetConf().JWT.SecretKey)
	})
}
