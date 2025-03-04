package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

func TestListProducts(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建多个测试商品
	var testProducts []*model.Product
	for i := 0; i < 5; i++ {
		product := createTestProduct(t, int64(category.ID))
		testProducts = append(testProducts, product)
		// 为每个商品创建一个SKU
		createTestSKU(t, product.ID)
	}

	t.Run("success with pagination", func(t *testing.T) {
		svc := NewListProductsService(getTestContext())

		// 测试第一页
		req := &product.ListProductsReq{
			PageSize: 2,
			Page:     1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list products: %v", err)
		}

		// 验证响应
		if resp.Total != 5 {
			t.Errorf("total products = %d, want 5", resp.Total)
		}
		if len(resp.Products) != 2 {
			t.Errorf("got %d products, want 2", len(resp.Products))
		}

		// 测试第二页
		req.Page = 2
		resp, err = svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list products page 2: %v", err)
		}

		if len(resp.Products) != 2 {
			t.Errorf("got %d products on page 2, want 2", len(resp.Products))
		}

		// 测试最后一页
		req.Page = 3
		resp, err = svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list products last page: %v", err)
		}

		if len(resp.Products) != 1 {
			t.Errorf("got %d products on last page, want 1", len(resp.Products))
		}
	})

	t.Run("empty list", func(t *testing.T) {
		// 清空商品表
		if err := mysql.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Product{}).Error; err != nil {
			t.Fatalf("failed to clear products: %v", err)
		}

		svc := NewListProductsService(getTestContext())

		req := &product.ListProductsReq{
			PageSize: 10,
			Page:     1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list products: %v", err)
		}

		if resp.Total != 0 {
			t.Errorf("total products = %d, want 0", resp.Total)
		}
		if len(resp.Products) != 0 {
			t.Errorf("got %d products, want 0", len(resp.Products))
		}
	})

	t.Run("invalid page parameters", func(t *testing.T) {
		svc := NewListProductsService(getTestContext())

		invalidCases := []struct {
			name     string
			page     int32
			pageSize int32
		}{
			{"zero page", 0, 10},
			{"negative page", -1, 10},
			{"zero page size", 1, 0},
			{"negative page size", 1, -1},
		}

		for _, tc := range invalidCases {
			t.Run(tc.name, func(t *testing.T) {
				req := &product.ListProductsReq{
					Page:     tc.page,
					PageSize: tc.pageSize,
				}

				resp, err := svc.Run(req)
				if err == nil {
					t.Error("expected error for invalid parameters, got nil")
				}
				if resp != nil && len(resp.Products) > 0 {
					t.Error("expected empty product list for invalid parameters")
				}
			})
		}
	})
}
