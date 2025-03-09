package service

import (
	"context"
	"encoding/json"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp = &product.CreateProductResp{}

	// 将商品分类转换为JSON字符串
	categories, err := json.Marshal(req.Categories)
	if err != nil {
		return nil, err
	}

	// 创建商品记录
	productModel := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  string(categories),
	}

	// 保存到数据库
	if err := model.CreateProduct(mysql.DB, productModel); err != nil {
		return nil, err
	}

	// 构建响应
	resp.Product = &product.Product{
		Id:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Picture:     productModel.Picture,
		Price:       productModel.Price,
		Categories:  req.Categories,
	}
	return resp, nil
}
