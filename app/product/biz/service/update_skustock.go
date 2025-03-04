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

type UpdateSKUStockService struct {
	ctx context.Context
}

// NewUpdateSKUStockService new UpdateSKUStockService
func NewUpdateSKUStockService(ctx context.Context) *UpdateSKUStockService {
	return &UpdateSKUStockService{ctx: ctx}
}

// Run update SKU stock
func (s *UpdateSKUStockService) Run(req *product.UpdateSKUStockReq) (resp *product.UpdateSKUStockResp, err error) {
	resp = new(product.UpdateSKUStockResp)

	// 验证库存值
	if req.Stock < 0 {
		err = errors.New("invalid stock quantity")
		klog.Errorf("%s: %d", err.Error(), req.Stock)
		return nil, err
	}

	maxRetries := 3
	currentTry := 0

	for currentTry < maxRetries {
		currentTry++

		// 开启事务
		err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
			// 1. 检查SKU是否存在并获取当前版本
			var sku model.SKU
			if err := tx.First(&sku, req.Id).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					klog.Errorf("sku not found: %v", err)
				} else {
					klog.Errorf("query sku failed: %v", err)
				}
				return err
			}

			// 2. 记录原始库存
			oldStock := sku.Stock
			currentVersion := sku.Version

			// 3. 更新库存（使用版本号作为条件）
			result := tx.Model(&sku).
				Where("id = ? AND version = ?", req.Id, currentVersion).
				Updates(map[string]interface{}{
					"stock":   req.Stock,
					"version": currentVersion + 1,
				})

			if result.Error != nil {
				klog.Errorf("update stock failed: %v", result.Error)
				return result.Error
			}

			if result.RowsAffected == 0 {
				klog.Warnf("stock update conflict detected, retry %d/%d", currentTry, maxRetries)
				return errors.New("stock update conflict")
			}

			// 4. 记录库存变更日志
			klog.Infof("SKU %d stock updated: %d -> %d (version: %d -> %d)",
				sku.ID, oldStock, req.Stock, currentVersion, currentVersion+1)

			// 5. 查询更新后的SKU信息
			var updatedSKU model.SKU
			if err := tx.First(&updatedSKU, req.Id).Error; err != nil {
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

		// 如果更新成功或遇到非冲突错误，直接返回
		if err == nil || !errors.Is(err, errors.New("stock update conflict")) {
			return resp, err
		}

		// 如果是最后一次重试失败
		if currentTry == maxRetries {
			klog.Errorf("failed to update stock after %d retries", maxRetries)
			return nil, errors.New("failed to update stock due to conflicts")
		}
	}

	return resp, nil
}
