package service

import (
	"encoding/json"

	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

// convertModelToProduct 将数据库模型转换为响应模型
func convertModelToProduct(p *model.Product) (*product.Product, error) {
	// 解析图片
	var images []string
	if err := json.Unmarshal([]byte(p.Images), &images); err != nil {
		return nil, err
	}

	// 转换SKUs
	var skus []*product.SKU
	for _, sku := range p.SKUs {
		var specs map[string]string
		if err := json.Unmarshal([]byte(sku.Specs), &specs); err != nil {
			return nil, err
		}

		skus = append(skus, &product.SKU{
			Id:        int64(sku.ID),
			ProductId: int64(sku.ProductID),
			Specs:     specs,
			Price:     sku.Price,
			Stock:     sku.Stock,
			Code:      sku.Code,
		})
	}

	return &product.Product{
		Id:          int64(p.ID),
		Name:        p.Name,
		Description: p.Description,
		CategoryId:  p.CategoryID,
		Images:      images,
		Price:       p.Price,
		Status:      p.Status,
		Skus:        skus,
		CreatedAt:   p.CreatedAt.Unix(),
		UpdatedAt:   p.UpdatedAt.Unix(),
	}, nil
}
