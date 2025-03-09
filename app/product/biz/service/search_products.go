package service

import (
"context"
"encoding/json"

"github.com/beatpika/eshop/app/product/biz/dal/mysql"
product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
ctx context.Context
}

// NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
return &SearchProductsService{ctx: ctx}
}

// Run search products
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
resp = &product.SearchProductsResp{}

// 计算分页参数
offset := (req.Page - 1) * int32(req.PageSize)
limit := int(req.PageSize)

// 搜索商品
products, total, err := mysql.SearchProducts(mysql.DB, req.Keywords, int(offset), limit)
if err != nil {
return nil, err
}

// 构建响应
resp.Results = make([]*product.Product, 0, len(products))
for _, p := range products {
var categories []string
if err := json.Unmarshal([]byte(p.Categories), &categories); err != nil {
return nil, err
}

resp.Results = append(resp.Results, &product.Product{
Id:          p.ID,
Name:        p.Name,
Description: p.Description,
Picture:     p.Picture,
Price:       p.Price,
Categories:  categories,
})
}

resp.Total = int32(total)
return resp, nil
}
