package service

import (
	"context"

	"github.com/beatpika/eshop/app/cart/biz/dal/redis"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

func GetCart(ctx context.Context, req *cart.GetCartReq) (*cart.GetCartResp, error) {
	// 从Redis获取购物车
	items, err := redis.GetCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	cartItems := make([]*cart.CartItem, 0, len(items))
	for _, item := range items {
		cartItems = append(cartItems, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  cartItems,
		},
	}, nil
}
