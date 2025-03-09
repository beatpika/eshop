package service

import (
"context"
"encoding/json"

"github.com/beatpika/eshop/app/product/biz/dal/mysql"
product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run get product info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp = &product.GetProductResp{}

	// 从数据库获取商品
	productModel, err := mysql.GetProductByID(mysql.DB, req.Id)
	if err != nil {
		return nil, err
	}

	// 解析商品分类
	var categories []string
	if err := json.Unmarshal([]byte(productModel.Categories), &categories); err != nil {
		return nil, err
	}

	// 构建响应
	resp.Product = &product.Product{
		Id:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Picture:     productModel.Picture,
		Price:       productModel.Price,
		Categories:  categories,
	}

	return resp, nil
}
