package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type GetProductsByCategoryService struct {
	ctx context.Context
} // NewGetProductsByCategoryService new GetProductsByCategoryService
func NewGetProductsByCategoryService(ctx context.Context) *GetProductsByCategoryService {
	return &GetProductsByCategoryService{ctx: ctx}
}

// Run create note info
func (s *GetProductsByCategoryService) Run(req *product.GetProductsByCategoryReq) (resp *product.GetProductsByCategoryResp, err error) {
	// Finish your business logic.

	return
}
