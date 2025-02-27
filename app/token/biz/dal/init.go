package dal

import (
	"github.com/beatpika/eshop/app/token/biz/dal/mysql"
	"github.com/beatpika/eshop/app/token/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
