package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type UpdateProductStatusService struct {
	ctx context.Context
} // NewUpdateProductStatusService new UpdateProductStatusService
func NewUpdateProductStatusService(ctx context.Context) *UpdateProductStatusService {
	return &UpdateProductStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductStatusService) Run(req *product.UpdateProductStatusReq) (resp *product.UpdateProductStatusResp, err error) {
	// Finish your business logic.

	return
}
