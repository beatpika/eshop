package service

import (
	"context"
	"fmt"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type DeleteSKUService struct {
	ctx context.Context
}

// NewDeleteSKUService new DeleteSKUService
func NewDeleteSKUService(ctx context.Context) *DeleteSKUService {
	return &DeleteSKUService{ctx: ctx}
}

// Run delete SKU
func (s *DeleteSKUService) Run(req *product.DeleteSKUReq) (resp *product.DeleteSKUResp, err error) {
	resp = new(product.DeleteSKUResp)

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查SKU是否存在，并锁定记录
		var sku model.SKU
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&sku, req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("sku not found: %v", err)
			} else {
				klog.Errorf("query sku failed: %v", err)
			}
			return err
		}

		// 2. 检查库存是否为0
		if sku.Stock > 0 {
			err := fmt.Errorf("cannot delete sku with stock: %d", sku.Stock)
			klog.Error(err)
			return err
		}

		// 3. 删除SKU（软删除）
		if err := tx.Delete(&sku).Error; err != nil {
			klog.Errorf("delete sku failed: %v", err)
			return err
		}

		// 4. 清除任何关联的缓存（如果有）
		// TODO: 实现缓存清理逻辑

		resp.Success = true
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
