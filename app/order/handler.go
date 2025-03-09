package main

import (
	"context"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
	"github.com/beatpika/eshop/app/order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusReq) (resp *order.UpdateOrderStatusResp, err error) {
	resp, err = service.NewUpdateOrderStatusService(ctx).Run(req)

	return resp, err
}
