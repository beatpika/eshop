package service

import (
	"context"
	"encoding/json"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
}

// NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run list products
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp = &product.ListProductsResp{}

	// 计算分页参数
	offset := (req.Page - 1) * int32(req.PageSize)
	limit := int(req.PageSize)

	// 从数据库获取商品列表
	products, total, err := model.ListProducts(mysql.DB, int(offset), limit, req.Category)
	if err != nil {
		return nil, err
	}

	// 构建响应
	resp.Products = make([]*product.Product, 0, len(products))
	for _, p := range products {
		var categories []string
		if err := json.Unmarshal([]byte(p.Categories), &categories); err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product.Product{
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
