package dal

import (
	"github.com/beatpika/eshop/app/order/biz/dal/mysql"
	"github.com/beatpika/eshop/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
