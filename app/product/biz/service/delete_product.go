package service

import (
	"context"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type DeleteProductService struct {
	ctx context.Context
}

// NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run delete product
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp = new(product.DeleteProductResp)

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查商品是否存在
		var prod model.Product
		if err := tx.First(&prod, req.Id).Error; err != nil {
			return err
		}

		// 2. 删除关联的SKUs（软删除）
		if err := tx.Where("product_id = ?", req.Id).Delete(&model.SKU{}).Error; err != nil {
			return err
		}

		// 3. 删除商品（软删除）
		if err := tx.Delete(&prod).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			klog.Errorf("product not found: %v", err)
		} else {
			klog.Errorf("delete product failed: %v", err)
		}
		return nil, err
	}

	resp.Success = true
	return resp, nil
}
