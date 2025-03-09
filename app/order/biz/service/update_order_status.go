package service

import (
	"context"
	order "github.com/beatpika/eshop/rpc_gen/kitex_gen/order"
)

type UpdateOrderStatusService struct {
	ctx context.Context
} // NewUpdateOrderStatusService new UpdateOrderStatusService
func NewUpdateOrderStatusService(ctx context.Context) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderStatusService) Run(req *order.UpdateOrderStatusReq) (resp *order.UpdateOrderStatusResp, err error) {
	// Finish your business logic.

	return
}
