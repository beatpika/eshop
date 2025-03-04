package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type CreateCategoryService struct {
	ctx context.Context
} // NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{ctx: ctx}
}

// Run create note info
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	// Finish your business logic.

	return
}
