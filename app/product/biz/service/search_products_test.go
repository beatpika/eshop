package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestSearchProducts(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category1 := createTestCategory(t, "手机", 0, 1, 1)
	category2 := createTestCategory(t, "电脑", 0, 1, 2)

	// 创建测试商品
	phones := []*model.Product{
		createTestProduct(t, int64(category1.ID)), // iPhone, 1000
		createTestProduct(t, int64(category1.ID)), // Huawei, 800
		createTestProduct(t, int64(category1.ID)), // Samsung, 900
	}

	// 设置不同的价格和名称
	mysql.DB.Model(phones[0]).Updates(map[string]interface{}{
		"name":  "iPhone手机",
		"price": 1000,
	})
	mysql.DB.Model(phones[1]).Updates(map[string]interface{}{
		"name":  "华为手机",
		"price": 800,
	})
	mysql.DB.Model(phones[2]).Updates(map[string]interface{}{
		"name":  "三星手机",
		"price": 900,
	})

	// 创建电脑类商品
	laptops := []*model.Product{
		createTestProduct(t, int64(category2.ID)),
	}
	mysql.DB.Model(laptops[0]).Updates(map[string]interface{}{
		"name":  "ThinkPad笔记本",
		"price": 2000,
	})

	t.Run("search by keyword", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			Keyword:  "手机",
			PageSize: 10,
			Page:     1,
			SortBy:   1, // 默认排序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products: %v", err)
		}

		// 验证结果
		if len(resp.Products) != 3 {
			t.Errorf("got %d products, want 3", len(resp.Products))
		}
		if resp.Total != 3 {
			t.Errorf("total = %d, want 3", resp.Total)
		}
	})

	t.Run("search by category", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			CategoryId: int64(category1.ID),
			PageSize:   10,
			Page:       1,
			SortBy:     1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products: %v", err)
		}

		// 验证结果
		if len(resp.Products) != 3 {
			t.Errorf("got %d products, want 3", len(resp.Products))
		}
		for _, p := range resp.Products {
			if p.CategoryId != int64(category1.ID) {
				t.Errorf("product category_id = %d, want %d", p.CategoryId, category1.ID)
			}
		}
	})

	t.Run("sort by price asc", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			Keyword:  "手机",
			PageSize: 10,
			Page:     1,
			SortBy:   2, // 价格升序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products: %v", err)
		}

		// 验证价格排序
		for i := 0; i < len(resp.Products)-1; i++ {
			if resp.Products[i].Price > resp.Products[i+1].Price {
				t.Errorf("products not properly sorted by price (asc)")
			}
		}
	})

	t.Run("sort by price desc", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			Keyword:  "手机",
			PageSize: 10,
			Page:     1,
			SortBy:   3, // 价格降序
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products: %v", err)
		}

		// 验证价格排序
		for i := 0; i < len(resp.Products)-1; i++ {
			if resp.Products[i].Price < resp.Products[i+1].Price {
				t.Errorf("products not properly sorted by price (desc)")
			}
		}
	})

	t.Run("pagination", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			Keyword:  "手机",
			PageSize: 2,
			Page:     1,
			SortBy:   1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products page 1: %v", err)
		}

		// 验证第一页
		if len(resp.Products) != 2 {
			t.Errorf("got %d products on page 1, want 2", len(resp.Products))
		}

		// 验证第二页
		req.Page = 2
		resp, err = svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products page 2: %v", err)
		}

		if len(resp.Products) != 1 {
			t.Errorf("got %d products on page 2, want 1", len(resp.Products))
		}
	})

	t.Run("no results", func(t *testing.T) {
		svc := NewSearchProductsService(getTestContext())

		req := &product.SearchProductsReq{
			Keyword:  "不存在的商品",
			PageSize: 10,
			Page:     1,
			SortBy:   1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to search products: %v", err)
		}

		if len(resp.Products) != 0 {
			t.Errorf("got %d products, want 0", len(resp.Products))
		}
		if resp.Total != 0 {
			t.Errorf("total = %d, want 0", resp.Total)
		}
	})
}
