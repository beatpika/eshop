package service

import (
	"context"

	"github.com/beatpika/eshop/app/cart/biz/dal/redis"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func AddItem(ctx context.Context, req *cart.AddItemReq) (*cart.AddItemResp, error) {
	// 创建购物车条目
	item := &redis.CartItem{
		ProductID: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	}

	// 添加到Redis
	if err := redis.AddItem(ctx, req.UserId, item); err != nil {
		return nil, err
	}

	return &cart.AddItemResp{}, nil
}
