package service

import (
	"context"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type ListCategoriesService struct {
	ctx context.Context
} // NewListCategoriesService new ListCategoriesService
func NewListCategoriesService(ctx context.Context) *ListCategoriesService {
	return &ListCategoriesService{ctx: ctx}
}

// Run create note info
func (s *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	// Finish your business logic.

	return
}
