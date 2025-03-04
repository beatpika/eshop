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

type UpdateProductService struct {
	ctx context.Context
}

// NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run update product
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp = new(product.UpdateProductResp)

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

		// 2. 检查商品是否存在
		var existingProduct model.Product
		if err := tx.First(&existingProduct, req.Id).Error; err != nil {
			klog.Errorf("update product failed: %v", err)
			return err
		}

		// 3. 转换图片列表为JSON
		imagesJSON, err := json.Marshal(req.Images)
		if err != nil {
			klog.Errorf("marshal images failed: %v", err)
			return err
		}

		// 4. 更新商品信息
		updates := map[string]interface{}{
			"name":        req.Name,
			"description": req.Description,
			"category_id": req.CategoryId,
			"images":      string(imagesJSON),
			"price":       req.Price,
		}

		if err := tx.Model(&existingProduct).Updates(updates).Error; err != nil {
			klog.Errorf("update product failed: %v", err)
			return err
		}

		// 5. 查询更新后的商品（包含SKUs）
		var updatedProduct model.Product
		if err := tx.Preload("SKUs").First(&updatedProduct, req.Id).Error; err != nil {
			return err
		}

		// 6. 转换为响应格式
		productResp, err := convertModelToProduct(&updatedProduct)
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
