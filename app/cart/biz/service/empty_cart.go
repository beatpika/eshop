package service

import (
	"context"

	"github.com/beatpika/eshop/app/cart/biz/dal/redis"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	// 清空购物车
	if err := redis.EmptyCart(ctx, req.UserId); err != nil {
		return nil, err
	}

	return &cart.EmptyCartResp{}, nil
}
