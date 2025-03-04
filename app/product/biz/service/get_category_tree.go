package service

import (
	"context"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCategoryTreeService struct {
	ctx context.Context
}

// NewGetCategoryTreeService new GetCategoryTreeService
func NewGetCategoryTreeService(ctx context.Context) *GetCategoryTreeService {
	return &GetCategoryTreeService{ctx: ctx}
}

// Run get category tree
func (s *GetCategoryTreeService) Run(req *product.GetCategoryTreeReq) (resp *product.GetCategoryTreeResp, err error) {
	resp = new(product.GetCategoryTreeResp)

	// 查询所有分类，按层级和排序序号排序
	var categories []*model.Category
	err = mysql.DB.WithContext(s.ctx).
		Order("level ASC, sort_order ASC").
		Find(&categories).Error
	if err != nil {
		klog.Errorf("get category tree failed: %v", err)
		return nil, err
	}

	// 转换为响应格式
	var categoryList []*product.Category
	for _, category := range categories {
		categoryList = append(categoryList, &product.Category{
			Id:        int64(category.ID),
			Name:      category.Name,
			ParentId:  category.ParentID,
			Level:     category.Level,
			SortOrder: category.SortOrder,
		})
	}

	resp.Categories = categoryList
	return resp, nil
}
