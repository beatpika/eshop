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

type UpdateSKUService struct {
	ctx context.Context
}

// NewUpdateSKUService new UpdateSKUService
func NewUpdateSKUService(ctx context.Context) *UpdateSKUService {
	return &UpdateSKUService{ctx: ctx}
}

// Run update SKU
func (s *UpdateSKUService) Run(req *product.UpdateSKUReq) (resp *product.UpdateSKUResp, err error) {
	resp = new(product.UpdateSKUResp)

	// 验证价格
	if req.Price < 0 {
		err = errors.New("price cannot be negative")
		klog.Errorf("invalid price: %d", req.Price)
		return nil, err
	}

	// 验证规格
	for key, value := range req.Specs {
		if key == "" || value == "" {
			err = errors.New("spec key and value cannot be empty")
			klog.Errorf("invalid specs: key=%s, value=%s", key, value)
			return nil, err
		}
	}

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查SKU是否存在
		var sku model.SKU
		if err := tx.First(&sku, req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("sku not found: %v", err)
			} else {
				klog.Errorf("query sku failed: %v", err)
			}
			return err
		}

		// 2. 检查新的SKU编码是否与其他SKU冲突
		if req.Code != sku.Code {
			var count int64
			if err := tx.Model(&model.SKU{}).Where("code = ? AND id != ?", req.Code, req.Id).Count(&count).Error; err != nil {
				return err
			}
			if count > 0 {
				err := errors.New("sku code already exists")
				klog.Errorf("%s: %s", err.Error(), req.Code)
				return err
			}
		}

		// 3. 将规格转换为JSON
		specsJSON, err := json.Marshal(req.Specs)
		if err != nil {
			klog.Errorf("marshal specs failed: %v", err)
			return err
		}

		// 4. 更新SKU信息
		sku.Specs = string(specsJSON)
		sku.Price = req.Price
		sku.Code = req.Code

		if err := tx.Save(&sku).Error; err != nil {
			klog.Errorf("update sku failed: %v", err)
			return err
		}

		// 5. 查询更新后的SKU信息
		var updatedSKU model.SKU
		if err := tx.First(&updatedSKU, sku.ID).Error; err != nil {
			return err
		}

		// 6. 转换为响应格式
		var specs map[string]string
		if err := json.Unmarshal([]byte(updatedSKU.Specs), &specs); err != nil {
			return err
		}

		resp.Sku = &product.SKU{
			Id:        int64(updatedSKU.ID),
			ProductId: int64(updatedSKU.ProductID),
			Specs:     specs,
			Price:     updatedSKU.Price,
			Stock:     updatedSKU.Stock,
			Code:      updatedSKU.Code,
		}

		return nil
	})
	if err != nil {
		klog.Errorf("update sku failed: %v", err)
		return nil, err
	}

	return resp, nil
}
