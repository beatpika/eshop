package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type UpdateSKUStockService struct {
	ctx context.Context
} // NewUpdateSKUStockService new UpdateSKUStockService
func NewUpdateSKUStockService(ctx context.Context) *UpdateSKUStockService {
	return &UpdateSKUStockService{ctx: ctx}
}

// Run create note info
func (s *UpdateSKUStockService) Run(req *product.UpdateSKUStockReq) (resp *product.UpdateSKUStockResp, err error) {
	// Finish your business logic.

	return
}
