package dal

import (
	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
