package dal

import (
	"github.com/beatpika/eshop/app/api/biz/dal/mysql"
	"github.com/beatpika/eshop/app/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
