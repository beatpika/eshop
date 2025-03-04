package service

import (
	"context"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListCategoriesService struct {
	ctx context.Context
}

// NewListCategoriesService new ListCategoriesService
func NewListCategoriesService(ctx context.Context) *ListCategoriesService {
	return &ListCategoriesService{ctx: ctx}
}

// Run list categories
func (s *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	resp = new(product.ListCategoriesResp)

	// 初始化空切片，确保在没有数据时返回空切片而不是nil
	resp.Categories = make([]*product.Category, 0)

	// 查询指定父分类下的所有子分类
	var categories []*model.Category
	err = mysql.DB.WithContext(s.ctx).
		Where("parent_id = ?", req.ParentId).
		Order("sort_order ASC").
		Find(&categories).Error
	if err != nil {
		klog.Errorf("list categories failed: %v", err)
		return nil, err
	}

	// 转换为响应格式
	for _, category := range categories {
		resp.Categories = append(resp.Categories, &product.Category{
			Id:        int64(category.ID),
			Name:      category.Name,
			ParentId:  category.ParentID,
			Level:     category.Level,
			SortOrder: category.SortOrder,
		})
	}

	return resp, nil
}
