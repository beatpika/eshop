package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateProduct(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)
	newCategory := createTestCategory(t, "New Category", 0, 1, 2)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	t.Run("success", func(t *testing.T) {
		svc := NewUpdateProductService(getTestContext())

		req := &product.UpdateProductReq{
			Id:          int64(testProduct.ID),
			Name:        "Updated Product",
			Description: "Updated Description",
			CategoryId:  int64(newCategory.ID),
			Images:      []string{"http://example.com/new-image.jpg"},
			Price:       2000, // 20.00
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update product: %v", err)
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
		if resp.Product.Price != req.Price {
			t.Errorf("product price = %d, want %d", resp.Product.Price, req.Price)
		}

		// 验证数据库记录
		var updatedProduct model.Product
		if err := mysql.DB.First(&updatedProduct, testProduct.ID).Error; err != nil {
			t.Fatalf("failed to query updated product: %v", err)
		}
		if updatedProduct.Name != req.Name {
			t.Errorf("db product name = %s, want %s", updatedProduct.Name, req.Name)
		}
		if updatedProduct.CategoryID != req.CategoryId {
			t.Errorf("db product category id = %d, want %d", updatedProduct.CategoryID, req.CategoryId)
		}
	})

	t.Run("product not found", func(t *testing.T) {
		svc := NewUpdateProductService(getTestContext())

		req := &product.UpdateProductReq{
			Id:          99999, // 不存在的商品ID
			Name:        "Updated Product",
			Description: "Updated Description",
			CategoryId:  int64(category.ID),
			Images:      []string{"http://example.com/new-image.jpg"},
			Price:       2000,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for non-existent product")
		}
	})

	t.Run("invalid category", func(t *testing.T) {
		svc := NewUpdateProductService(getTestContext())

		req := &product.UpdateProductReq{
			Id:          int64(testProduct.ID),
			Name:        "Updated Product",
			Description: "Updated Description",
			CategoryId:  99999, // 不存在的分类ID
			Images:      []string{"http://example.com/new-image.jpg"},
			Price:       2000,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for invalid category, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for invalid category")
		}
	})
}
