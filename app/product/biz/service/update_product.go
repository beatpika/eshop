package service

import (
"context"
"encoding/json"

"github.com/beatpika/eshop/app/product/biz/dal/mysql"
product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type UpdateProductService struct {
ctx context.Context
}

// NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
return &UpdateProductService{ctx: ctx}
}

// Run update product info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
resp = &product.UpdateProductResp{}

// 检查商品是否存在
productModel, err := mysql.GetProductByID(mysql.DB, req.Id)
if err != nil {
return nil, err
}

// 将商品分类转换为JSON字符串
categories, err := json.Marshal(req.Categories)
if err != nil {
return nil, err
}

// 更新商品信息
productModel.Name = req.Name
productModel.Description = req.Description
productModel.Picture = req.Picture
productModel.Price = req.Price
productModel.Categories = string(categories)

// 保存到数据库
if err := mysql.UpdateProduct(mysql.DB, productModel); err != nil {
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
