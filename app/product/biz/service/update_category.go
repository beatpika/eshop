package service

import (
	"context"
	"errors"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type UpdateCategoryService struct {
	ctx context.Context
}

// NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{ctx: ctx}
}

// Run update category
func (s *UpdateCategoryService) Run(req *product.UpdateCategoryReq) (resp *product.UpdateCategoryResp, err error) {
	resp = new(product.UpdateCategoryResp)

	// 验证名称不能为空
	if req.Name == "" {
		err = errors.New("category name cannot be empty")
		klog.Errorf("%s: %s", err.Error(), req.Name)
		return nil, err
	}

	// 验证排序序号不能为负数
	if req.SortOrder < 0 {
		err = errors.New("sort order cannot be negative")
		klog.Errorf("%s: %d", err.Error(), req.SortOrder)
		return nil, err
	}

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

		// 2. 更新分类信息（只更新名称和排序序号）
		category.Name = req.Name
		category.SortOrder = req.SortOrder

		if err := tx.Save(&category).Error; err != nil {
			klog.Errorf("update category failed: %v", err)
			return err
		}

		// 3. 查询更新后的分类信息
		var updatedCategory model.Category
		if err := tx.First(&updatedCategory, req.Id).Error; err != nil {
			return err
		}

		// 4. 转换为响应格式
		resp.Category = &product.Category{
			Id:        int64(updatedCategory.ID),
			Name:      updatedCategory.Name,
			ParentId:  updatedCategory.ParentID,
			Level:     updatedCategory.Level,
			SortOrder: updatedCategory.SortOrder,
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
