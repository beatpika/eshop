package service

import (
	"context"
	"errors"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListProductsService struct {
	ctx context.Context
}

// NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run list products
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp = new(product.ListProductsResp)
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

	// 初始化查询
	query := mysql.DB.WithContext(s.ctx).Model(&model.Product{})

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		klog.Errorf("count products failed: %v", err)
		return nil, err
	}
	resp.Total = int32(total)

	// 添加分页
	offset := (req.Page - 1) * req.PageSize
	query = query.Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(req.PageSize))

	// 执行查询
	var products []*model.Product
	if err := query.Preload("SKUs").Find(&products).Error; err != nil {
		klog.Errorf("list products failed: %v", err)
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
