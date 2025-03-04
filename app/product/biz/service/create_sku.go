package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type CreateSKUService struct {
	ctx context.Context
} // NewCreateSKUService new CreateSKUService
func NewCreateSKUService(ctx context.Context) *CreateSKUService {
	return &CreateSKUService{ctx: ctx}
}

// Run create note info
func (s *CreateSKUService) Run(req *product.CreateSKUReq) (resp *product.CreateSKUResp, err error) {
	// Finish your business logic.

	return
}
