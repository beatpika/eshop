package service

import (
"context"

"github.com/beatpika/eshop/app/product/biz/dal/mysql"
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
_, err = mysql.GetProductByID(mysql.DB, req.Id)
if err != nil {
return nil, err
}

// 从数据库删除商品
if err := mysql.DeleteProduct(mysql.DB, req.Id); err != nil {
return nil, err
}

// 构建响应
resp.Id = req.Id

return resp, nil
}
