package service

import (
	"context"
	"encoding/json"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetProductService struct {
	ctx context.Context
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run get product
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp = new(product.GetProductResp)

	var prod model.Product
	if err := mysql.DB.WithContext(s.ctx).Preload("SKUs").Preload("Category").First(&prod, req.Id).Error; err != nil {
		klog.Errorf("get product failed: %v", err)
		return nil, err
	}

	// 转换SKU信息
	var skus []*product.SKU
	for _, sku := range prod.SKUs {
		var specs map[string]string
		if err := json.Unmarshal([]byte(sku.Specs), &specs); err != nil {
			klog.Errorf("unmarshal sku specs failed: %v", err)
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

	// 转换图片列表
	var images []string
	if err := json.Unmarshal([]byte(prod.Images), &images); err != nil {
		klog.Errorf("unmarshal product images failed: %v", err)
		return nil, err
	}

	resp.Product = &product.Product{
		Id:          int64(prod.ID),
		Name:        prod.Name,
		Description: prod.Description,
		CategoryId:  prod.CategoryID,
		Images:      images,
		Price:       prod.Price,
		Status:      prod.Status,
		Skus:        skus,
		CreatedAt:   prod.CreatedAt.Unix(),
		UpdatedAt:   prod.UpdatedAt.Unix(),
	}

	return resp, nil
}
