package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestGetProductsByCategory(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "手机", 0, 1, 1)
	emptyCategory := createTestCategory(t, "空分类", 0, 1, 2)

	// 创建测试商品
	phones := []*model.Product{
		createTestProduct(t, int64(category.ID)),
		createTestProduct(t, int64(category.ID)),
		createTestProduct(t, int64(category.ID)),
		createTestProduct(t, int64(category.ID)),
	}

	// 设置不同的价格和名称
	priceSettings := []struct {
		name  string
		price int64
	}{
		{"高端手机", 2000},
		{"中端手机", 1500},
		{"入门手机", 1000},
		{"经济手机", 500},
	}

	for i, setting := range priceSettings {
		mysql.DB.Model(phones[i]).Updates(map[string]interface{}{
			"name":  setting.name,
			"price": setting.price,
		})
	}

	t.Run("list all products in category", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		req := &product.GetProductsByCategoryReq{
			CategoryId: int64(category.ID),
			PageSize:   10,
			Page:       1,
			SortBy:     1, // 默认排序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products: %v", err)
		}

		// 验证结果
		if len(resp.Products) != 4 {
			t.Errorf("got %d products, want 4", len(resp.Products))
		}
		if resp.Total != 4 {
			t.Errorf("total = %d, want 4", resp.Total)
		}

		// 验证所有商品都属于正确的分类
		for _, p := range resp.Products {
			if p.CategoryId != int64(category.ID) {
				t.Errorf("product category_id = %d, want %d", p.CategoryId, category.ID)
			}
		}
	})

	t.Run("sort by price ascending", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		req := &product.GetProductsByCategoryReq{
			CategoryId: int64(category.ID),
			PageSize:   10,
			Page:       1,
			SortBy:     2, // 价格升序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products: %v", err)
		}

		// 验证价格排序
		for i := 0; i < len(resp.Products)-1; i++ {
			if resp.Products[i].Price > resp.Products[i+1].Price {
				t.Errorf("products not properly sorted by price (asc)")
			}
		}
	})

	t.Run("sort by price descending", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		req := &product.GetProductsByCategoryReq{
			CategoryId: int64(category.ID),
			PageSize:   10,
			Page:       1,
			SortBy:     3, // 价格降序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products: %v", err)
		}

		// 验证价格排序
		for i := 0; i < len(resp.Products)-1; i++ {
			if resp.Products[i].Price < resp.Products[i+1].Price {
				t.Errorf("products not properly sorted by price (desc)")
			}
		}
	})

	t.Run("pagination", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		// 测试第一页
		req := &product.GetProductsByCategoryReq{
			CategoryId: int64(category.ID),
			PageSize:   2,
			Page:       1,
			SortBy:     1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products page 1: %v", err)
		}

		if len(resp.Products) != 2 {
			t.Errorf("got %d products on page 1, want 2", len(resp.Products))
		}

		// 测试第二页
		req.Page = 2
		resp, err = svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products page 2: %v", err)
		}

		if len(resp.Products) != 2 {
			t.Errorf("got %d products on page 2, want 2", len(resp.Products))
		}
	})

	t.Run("empty category", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		req := &product.GetProductsByCategoryReq{
			CategoryId: int64(emptyCategory.ID),
			PageSize:   10,
			Page:       1,
			SortBy:     1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get products: %v", err)
		}

		if len(resp.Products) != 0 {
			t.Errorf("got %d products, want 0", len(resp.Products))
		}
		if resp.Total != 0 {
			t.Errorf("total = %d, want 0", resp.Total)
		}
	})

	t.Run("non-existent category", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

		req := &product.GetProductsByCategoryReq{
			CategoryId: 99999, // 不存在的分类ID
			PageSize:   10,
			Page:       1,
			SortBy:     1,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent category, got nil")
		}
		if resp != nil && len(resp.Products) > 0 {
			t.Error("expected empty product list for non-existent category")
		}
	})

	t.Run("invalid page parameters", func(t *testing.T) {
		svc := NewGetProductsByCategoryService(getTestContext())

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
				req := &product.GetProductsByCategoryReq{
					CategoryId: int64(category.ID),
					Page:       tc.page,
					PageSize:   tc.pageSize,
					SortBy:     1,
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
