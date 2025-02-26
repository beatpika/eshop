package rpc

import (
	"sync"

	"github.com/beatpika/eshop/app/api/conf"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		InitUserClient()
	})
}

func InitUserClient() {
	// 使用 Consul 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		panic(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
