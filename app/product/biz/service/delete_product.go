package service

import (
	"context"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
}

// NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run delete product
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp = &product.DeleteProductResp{}

	// 检查商品是否存在
	_, err = model.GetProductByID(mysql.DB, req.Id)
	if err != nil {
		return nil, err
	}

	// 从数据库删除商品
	if err := model.DeleteProduct(mysql.DB, req.Id); err != nil {
		return nil, err
	}

	// 构建响应
	resp.Id = req.Id

	return resp, nil
}
