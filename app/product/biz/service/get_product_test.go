package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestGetProduct(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	// 创建测试SKU
	testSKU := createTestSKU(t, testProduct.ID)

	t.Run("success", func(t *testing.T) {
		svc := NewGetProductService(getTestContext())

		req := &product.GetProductReq{
			Id: int64(testProduct.ID),
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to get product: %v", err)
		}

		// 验证响应
		if resp.Product == nil {
			t.Fatal("product in response is nil")
		}

		// 验证商品基本信息
		if resp.Product.Id != int64(testProduct.ID) {
			t.Errorf("product id = %d, want %d", resp.Product.Id, testProduct.ID)
		}
		if resp.Product.Name != testProduct.Name {
			t.Errorf("product name = %s, want %s", resp.Product.Name, testProduct.Name)
		}
		if resp.Product.CategoryId != int64(testProduct.CategoryID) {
			t.Errorf("product category id = %d, want %d", resp.Product.CategoryId, testProduct.CategoryID)
		}

		// 验证SKU信息
		if len(resp.Product.Skus) != 1 {
			t.Errorf("got %d skus, want 1", len(resp.Product.Skus))
		} else {
			sku := resp.Product.Skus[0]
			if sku.Id != int64(testSKU.ID) {
				t.Errorf("sku id = %d, want %d", sku.Id, testSKU.ID)
			}
			if sku.ProductId != int64(testSKU.ProductID) {
				t.Errorf("sku product id = %d, want %d", sku.ProductId, testSKU.ProductID)
			}
			if sku.Price != testSKU.Price {
				t.Errorf("sku price = %d, want %d", sku.Price, testSKU.Price)
			}
			if sku.Stock != testSKU.Stock {
				t.Errorf("sku stock = %d, want %d", sku.Stock, testSKU.Stock)
			}
			if sku.Code != testSKU.Code {
				t.Errorf("sku code = %s, want %s", sku.Code, testSKU.Code)
			}
		}
	})

	t.Run("product not found", func(t *testing.T) {
		svc := NewGetProductService(getTestContext())

		req := &product.GetProductReq{
			Id: 99999, // 不存在的商品ID
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for non-existent product")
		}
	})

	t.Run("deleted product", func(t *testing.T) {
		// 软删除测试商品
		if err := mysql.DB.Delete(&testProduct).Error; err != nil {
			t.Fatalf("failed to delete test product: %v", err)
		}

		svc := NewGetProductService(getTestContext())

		req := &product.GetProductReq{
			Id: int64(testProduct.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for deleted product, got nil")
		}
		if resp != nil && resp.Product != nil {
			t.Error("expected nil product for deleted product")
		}
	})
}
