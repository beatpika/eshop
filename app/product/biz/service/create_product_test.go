package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestCreateProduct(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	t.Run("success", func(t *testing.T) {
		svc := NewCreateProductService(getTestContext())

		// 创建商品请求
		req := &product.CreateProductReq{
			Name:        "Test Product",
			Description: "Test Description",
			CategoryId:  int64(category.ID),
			Images:      []string{"http://example.com/image1.jpg", "http://example.com/image2.jpg"},
			Price:       1000, // 10.00
			Skus: []*product.SKU{
				{
					Specs: map[string]string{
						"color": "Red",
						"size":  "L",
					},
					Price: 1000,
					Stock: 100,
					Code:  "TST-001",
				},
				{
					Specs: map[string]string{
						"color": "Blue",
						"size":  "M",
					},
					Price: 1000,
					Stock: 50,
					Code:  "TST-002",
				},
			},
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to create product: %v", err)
		}

		// 验证响应
		if resp.Product == nil {
			t.Fatal("product in response is nil")
		}
		if resp.Product.Name != req.Name {
			t.Errorf("product name = %s, want %s", resp.Product.Name, req.Name)
		}
		if resp.Product.CategoryId != req.CategoryId {
			t.Errorf("product category id = %d, want %d", resp.Product.CategoryId, req.CategoryId)
		}
		if len(resp.Product.Skus) != len(req.Skus) {
			t.Errorf("got %d skus, want %d", len(resp.Product.Skus), len(req.Skus))
		}

		// 验证数据库记录
		var count int64
		if err := mysql.DB.Model(&model.Product{}).Where("id = ?", resp.Product.Id).Count(&count).Error; err != nil {
			t.Fatalf("failed to query product: %v", err)
		}
		if count != 1 {
			t.Errorf("product count = %d, want 1", count)
		}

		var skuCount int64
		if err := mysql.DB.Model(&model.SKU{}).Where("product_id = ?", resp.Product.Id).Count(&skuCount).Error; err != nil {
			t.Fatalf("failed to query skus: %v", err)
		}
		if skuCount != int64(len(req.Skus)) {
			t.Errorf("sku count = %d, want %d", skuCount, len(req.Skus))
		}
	})

	t.Run("invalid category", func(t *testing.T) {
		svc := NewCreateProductService(getTestContext())

		req := &product.CreateProductReq{
			Name:        "Test Product",
			Description: "Test Description",
			CategoryId:  999999, // 不存在的分类ID
			Images:      []string{"http://example.com/image1.jpg"},
			Price:       1000,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for invalid category, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for invalid category")
		}
	})

	t.Run("duplicate sku code", func(t *testing.T) {
		svc := NewCreateProductService(getTestContext())

		req := &product.CreateProductReq{
			Name:        "Test Product",
			Description: "Test Description",
			CategoryId:  int64(category.ID),
			Images:      []string{"http://example.com/image1.jpg"},
			Price:       1000,
			Skus: []*product.SKU{
				{
					Specs: map[string]string{"color": "Red"},
					Price: 1000,
					Stock: 100,
					Code:  "TST-001",
				},
				{
					Specs: map[string]string{"color": "Blue"},
					Price: 1000,
					Stock: 50,
					Code:  "TST-001", // 重复的SKU编码
				},
			},
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for duplicate sku code, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for duplicate sku code")
		}
	})
}
