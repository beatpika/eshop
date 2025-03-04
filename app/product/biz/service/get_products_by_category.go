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

type GetProductsByCategoryService struct {
	ctx context.Context
}

// NewGetProductsByCategoryService new GetProductsByCategoryService
func NewGetProductsByCategoryService(ctx context.Context) *GetProductsByCategoryService {
	return &GetProductsByCategoryService{ctx: ctx}
}

// Run get products by category
func (s *GetProductsByCategoryService) Run(req *product.GetProductsByCategoryReq) (resp *product.GetProductsByCategoryResp, err error) {
	resp = new(product.GetProductsByCategoryResp)

	// 初始化空切片，确保在没有数据时返回空切片而不是nil
	resp.Products = make([]*product.Product, 0)

	// 验证分页参数
	if req.Page <= 0 {
		err = errors.New("page number must be positive")
		klog.Errorf("%s: %d", err.Error(), req.Page)
		return nil, err
	}
	if req.PageSize <= 0 {
		err = errors.New("page size must be positive")
		klog.Errorf("%s: %d", err.Error(), req.PageSize)
		return nil, err
	}

	// 验证分类是否存在
	var category model.Category
	if err := mysql.DB.WithContext(s.ctx).First(&category, req.CategoryId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			klog.Errorf("category not found: %v", err)
		} else {
			klog.Errorf("query category failed: %v", err)
		}
		return nil, err
	}

	// 初始化查询
	query := mysql.DB.WithContext(s.ctx).Model(&model.Product{}).
		Where("category_id = ?", req.CategoryId)

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		klog.Errorf("count products failed: %v", err)
		return nil, err
	}
	resp.Total = int32(total)

	// 添加排序
	switch req.SortBy {
	case 2: // 价格升序
		query = query.Order("price ASC")
	case 3: // 价格降序
		query = query.Order("price DESC")
	case 4: // 销量排序
		query = query.Order("sales DESC")
	default: // 默认排序
		query = query.Order("created_at DESC")
	}

	// 添加分页
	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(int(offset)).Limit(int(req.PageSize))

	// 执行查询
	var products []*model.Product
	if err := query.Preload("SKUs").Find(&products).Error; err != nil {
		klog.Errorf("query products failed: %v", err)
		return nil, err
	}

	// 转换为响应格式
	for _, p := range products {
		productResp, err := convertModelToProduct(p)
		if err != nil {
			klog.Errorf("convert product failed: %v", err)
			return nil, err
		}
		resp.Products = append(resp.Products, productResp)
	}

	return resp, nil
}
