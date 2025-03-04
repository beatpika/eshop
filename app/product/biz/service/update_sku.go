package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type UpdateSKUService struct {
	ctx context.Context
} // NewUpdateSKUService new UpdateSKUService
func NewUpdateSKUService(ctx context.Context) *UpdateSKUService {
	return &UpdateSKUService{ctx: ctx}
}

// Run create note info
func (s *UpdateSKUService) Run(req *product.UpdateSKUReq) (resp *product.UpdateSKUResp, err error) {
	// Finish your business logic.

	return
}
