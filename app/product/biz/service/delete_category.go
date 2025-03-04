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

type DeleteCategoryService struct {
	ctx context.Context
}

// NewDeleteCategoryService new DeleteCategoryService
func NewDeleteCategoryService(ctx context.Context) *DeleteCategoryService {
	return &DeleteCategoryService{ctx: ctx}
}

// Run delete category
func (s *DeleteCategoryService) Run(req *product.DeleteCategoryReq) (resp *product.DeleteCategoryResp, err error) {
	resp = new(product.DeleteCategoryResp)

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查分类是否存在
		var category model.Category
		if err := tx.First(&category, req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				klog.Errorf("category not found: %v", err)
			} else {
				klog.Errorf("query category failed: %v", err)
			}
			return err
		}

		// 2. 检查是否有子分类
		var subCategoryCount int64
		if err := tx.Model(&model.Category{}).Where("parent_id = ?", req.Id).Count(&subCategoryCount).Error; err != nil {
			return err
		}
		if subCategoryCount > 0 {
			err := fmt.Errorf("category has %d sub-categories, cannot delete", subCategoryCount)
			klog.Error(err)
			return err
		}

		// 3. 检查是否有关联的商品
		var productCount int64
		if err := tx.Model(&model.Product{}).Where("category_id = ?", req.Id).Count(&productCount).Error; err != nil {
			return err
		}
		if productCount > 0 {
			err := fmt.Errorf("category has %d products, cannot delete", productCount)
			klog.Error(err)
			return err
		}

		// 4. 删除分类（软删除）
		if err := tx.Delete(&category).Error; err != nil {
			klog.Errorf("delete category failed: %v", err)
			return err
		}

		resp.Success = true
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
