package service

import (
	"context"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SearchProductsService struct {
	ctx context.Context
}

// NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run search products
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp = new(product.SearchProductsResp)

	// 验证分页参数
	if req.Page <= 0 || req.PageSize <= 0 {
		klog.Errorf("invalid page parameters: page=%d, page_size=%d", req.Page, req.PageSize)
		return nil, err
	}

	// 初始化查询
	query := mysql.DB.WithContext(s.ctx).Model(&model.Product{})

	// 添加分类筛选
	if req.CategoryId > 0 {
		query = query.Where("category_id = ?", req.CategoryId)
	}

	// 添加关键词搜索
	if req.Keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?",
			"%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

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
	case 4: // 销量排序（如果有销量字段）
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
		klog.Errorf("search products failed: %v", err)
		return nil, err
	}

	// 转换为响应格式
	resp.Products = make([]*product.Product, 0, len(products))
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
