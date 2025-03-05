package main

import (
	"context"

	"github.com/beatpika/eshop/app/cart/biz/service"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/cart"
)

// CartServiceImpl implements the cart service interface.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	return service.AddItem(ctx, req)
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	return service.GetCart(ctx, req)
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	return service.EmptyCart(ctx, req)
}
