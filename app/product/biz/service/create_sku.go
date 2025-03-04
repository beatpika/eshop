package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type CreateSKUService struct {
	ctx context.Context
}

// NewCreateSKUService new CreateSKUService
func NewCreateSKUService(ctx context.Context) *CreateSKUService {
	return &CreateSKUService{ctx: ctx}
}

// Run create SKU
func (s *CreateSKUService) Run(req *product.CreateSKUReq) (resp *product.CreateSKUResp, err error) {
	resp = new(product.CreateSKUResp)

	// 基础参数验证
	if req.Code == "" {
		err = errors.New("SKU code cannot be empty")
		klog.Error(err)
		return nil, err
	}

	if req.Price < 0 {
		err = errors.New("price cannot be negative")
		klog.Errorf("%s: %d", err.Error(), req.Price)
		return nil, err
	}

	if req.Stock < 0 {
		err = errors.New("stock cannot be negative")
		klog.Errorf("%s: %d", err.Error(), req.Stock)
		return nil, err
	}

	// 规格验证
	for key, value := range req.Specs {
		if key == "" || value == "" {
			err = fmt.Errorf("invalid specs: key=%s, value=%s", key, value)
			klog.Error(err)
			return nil, err
		}
	}

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 验证商品是否存在
		var existingProduct model.Product
		if err := tx.First(&existingProduct, req.ProductId).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("product not found: %v", err)
				return errors.New("product not found")
			}
			klog.Errorf("query product failed: %v", err)
			return err
		}

		// 2. 验证SKU编码唯一性
		var count int64
		if err := tx.Model(&model.SKU{}).Where("code = ?", req.Code).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			err := fmt.Errorf("SKU code already exists: %s", req.Code)
			klog.Error(err)
			return err
		}

		// 3. 将规格转换为JSON
		specsJSON, err := json.Marshal(req.Specs)
		if err != nil {
			klog.Errorf("marshal specs failed: %v", err)
			return err
		}

		// 4. 创建SKU
		sku := &model.SKU{
			ProductID: uint(req.ProductId),
			Specs:     string(specsJSON),
			Price:     req.Price,
			Stock:     req.Stock,
			Code:      req.Code,
			Version:   1, // 初始版本号
		}

		if err := tx.Create(sku).Error; err != nil {
			klog.Errorf("create sku failed: %v", err)
			return err
		}

		// 返回创建的SKU信息
		resp.Sku = &product.SKU{
			Id:        int64(sku.ID),
			ProductId: int64(sku.ProductID),
			Specs:     req.Specs,
			Price:     sku.Price,
			Stock:     sku.Stock,
			Code:      sku.Code,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
