package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type DeleteCategoryService struct {
	ctx context.Context
} // NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	// Finish your business logic.

	return
}
