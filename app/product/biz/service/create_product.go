package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type CreateProductService struct {
	ctx context.Context
}

// NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create product
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp = new(product.CreateProductResp)

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 验证分类是否存在
		var category model.Category
		if err := tx.First(&category, req.CategoryId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("category not found: %v", err)
				return errors.New("category not found")
			}
			klog.Errorf("query category failed: %v", err)
			return err
		}

		// 2. 验证SKU编码唯一性
		var existingSKUCount int64
		for _, sku := range req.Skus {
			if err := tx.Model(&model.SKU{}).Where("code = ?", sku.Code).Count(&existingSKUCount).Error; err != nil {
				return err
			}
			if existingSKUCount > 0 {
				klog.Errorf("duplicate sku code: %s", sku.Code)
				return errors.New("duplicate sku code")
			}
		}

		// 3. 转换图片列表为JSON
		imagesJSON, err := json.Marshal(req.Images)
		if err != nil {
			klog.Errorf("marshal images failed: %v", err)
			return err
		}

		// 4. 创建商品
		product := &model.Product{
			Name:        req.Name,
			Description: req.Description,
			CategoryID:  req.CategoryId,
			Images:      string(imagesJSON),
			Price:       req.Price,
			Status:      1, // 默认为待上架状态
		}

		if err := tx.Create(product).Error; err != nil {
			klog.Errorf("create product failed: %v", err)
			return err
		}

		// 5. 创建SKUs
		for _, skuReq := range req.Skus {
			specsJSON, err := json.Marshal(skuReq.Specs)
			if err != nil {
				klog.Errorf("marshal specs failed: %v", err)
				return err
			}

			sku := &model.SKU{
				ProductID: product.ID,
				Specs:     string(specsJSON),
				Price:     skuReq.Price,
				Stock:     skuReq.Stock,
				Code:      skuReq.Code,
				Version:   1, // 初始版本号
			}

			if err := tx.Create(sku).Error; err != nil {
				klog.Errorf("create sku failed: %v", err)
				return err
			}
		}

		// 6. 查询创建的商品（包含SKUs）
		var createdProduct model.Product
		if err := tx.Preload("SKUs").First(&createdProduct, product.ID).Error; err != nil {
			return err
		}

		// 7. 转换为响应格式
		productResp, err := convertModelToProduct(&createdProduct)
		if err != nil {
			return err
		}
		resp.Product = productResp

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
