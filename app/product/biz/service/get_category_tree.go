package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type GetCategoryTreeService struct {
	ctx context.Context
} // NewGetCategoryTreeService new GetCategoryTreeService
func NewGetCategoryTreeService(ctx context.Context) *GetCategoryTreeService {
	return &GetCategoryTreeService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryTreeService) Run(req *product.GetCategoryTreeReq) (resp *product.GetCategoryTreeResp, err error) {
	// Finish your business logic.

	return
}
