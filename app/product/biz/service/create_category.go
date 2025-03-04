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

type CreateCategoryService struct {
	ctx context.Context
}

// NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{ctx: ctx}
}

// Run create category
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	resp = new(product.CreateCategoryResp)

	// 1. 验证分类层级
	if req.Level < 1 || req.Level > 3 {
		klog.Errorf("invalid category level: %d", req.Level)
		err = fmt.Errorf("invalid category level: %d", req.Level)
		return nil, err
	}

	// 开启事务
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		// 2. 如果不是一级分类，检查父分类是否存在
		if req.ParentId != 0 {
			var parentCategory model.Category
			if err := tx.First(&parentCategory, req.ParentId).Error; err != nil {
				return err
			}

			// 检查父分类的层级是否合法
			if parentCategory.Level >= req.Level {
				klog.Errorf("invalid parent level: parent=%d, current=%d", parentCategory.Level, req.Level)
				err = fmt.Errorf("invalid parent level: parent=%d, current=%d", parentCategory.Level, req.Level)
				return err
			}
		}

		// 3. 创建分类
		category := &model.Category{
			Name:      req.Name,
			ParentID:  req.ParentId,
			Level:     req.Level,
			SortOrder: req.SortOrder,
		}

		if err := tx.Create(category).Error; err != nil {
			return err
		}

		// 4. 查询创建后的分类信息
		var createdCategory model.Category
		if err := tx.First(&createdCategory, category.ID).Error; err != nil {
			return err
		}

		// 5. 转换为响应格式
		resp.Category = &product.Category{
			Id:        int64(createdCategory.ID),
			Name:      createdCategory.Name,
			ParentId:  createdCategory.ParentID,
			Level:     createdCategory.Level,
			SortOrder: createdCategory.SortOrder,
		}

		return nil
	})
	if err != nil {
		klog.Errorf("create category failed: %v", err)
		return nil, err
	}

	return resp, nil
}
