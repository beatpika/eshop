package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UpdateProductStatusService struct {
	ctx context.Context
}

// NewUpdateProductStatusService new UpdateProductStatusService
func NewUpdateProductStatusService(ctx context.Context) *UpdateProductStatusService {
	return &UpdateProductStatusService{ctx: ctx}
}

// Run update product status
func (s *UpdateProductStatusService) Run(req *product.UpdateProductStatusReq) (resp *product.UpdateProductStatusResp, err error) {
	resp = new(product.UpdateProductStatusResp)

	// 验证状态值
	validStatuses := map[int32]bool{
		1: true, // 待上架
		2: true, // 已上架
		3: true, // 已下架
	}

	if !validStatuses[req.Status] {
		err = fmt.Errorf("invalid status: %d", req.Status)
		klog.Error(err)
		return nil, err
	}

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查商品是否存在
		var product model.Product
		if err := tx.First(&product, req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("product not found: %v", err)
				return errors.New("product not found")
			}
			klog.Errorf("query product failed: %v", err)
			return err
		}

		// 2. 更新商品状态
		product.Status = req.Status

		if err := tx.Save(&product).Error; err != nil {
			klog.Errorf("update product status failed: %v", err)
			return err
		}

		// 3. 查询更新后的商品信息
		var updatedProduct model.Product
		if err := tx.First(&updatedProduct, req.Id).Error; err != nil {
			return err
		}

		// 4. 转换为响应格式
		resp.Success = true

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
