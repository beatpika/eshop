package dal

import (
	"github.com/beatpika/eshop/app/cart/biz/dal/mysql"
	"github.com/beatpika/eshop/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
