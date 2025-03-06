package rpc

import (
	"fmt"
	"sync"

	"github.com/beatpika/eshop/app/api/conf"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth/authservice"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product/productservice"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient    userservice.Client
	TokenClient   authservice.Client
	ProductClient productservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		InitUserClient()
		InitTokenClient()
		InitProductClient()
		InitCartClient()
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
	fmt.Println("user client init success")
}

func InitTokenClient() {
	// 使用 Consul 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		panic(err)
	}
	TokenClient, err = authservice.NewClient("token", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func InitProductClient() {
	// 使用 Consul 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		panic(err)
	}
	ProductClient, err = productservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func InitCartClient() {
	// 使用 Consul 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		panic(err)
	}
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
