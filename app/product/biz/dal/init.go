package dal

import (
	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
