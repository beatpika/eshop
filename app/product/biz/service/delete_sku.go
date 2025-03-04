package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type DeleteSKUService struct {
	ctx context.Context
} // NewDeleteSKUService new DeleteSKUService
func NewDeleteSKUService(ctx context.Context) *DeleteSKUService {
	return &DeleteSKUService{ctx: ctx}
}

// Run create note info
func (s *DeleteSKUService) Run(req *product.DeleteSKUReq) (resp *product.DeleteSKUResp, err error) {
	// Finish your business logic.

	return
}
