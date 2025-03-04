package dal

import (
	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/dal/redis"
)

// Init init dal
func Init() {
	mysql.Init()
	redis.Init()
}
