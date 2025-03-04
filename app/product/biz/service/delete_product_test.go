package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

func TestDeleteProduct(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品和SKUs
	testProduct := createTestProduct(t, int64(category.ID))
	_ = createTestSKU(t, testProduct.ID)
	_ = createTestSKU(t, testProduct.ID)

	t.Run("success", func(t *testing.T) {
		svc := NewDeleteProductService(getTestContext())

		req := &product.DeleteProductReq{
			Id: int64(testProduct.ID),
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to delete product: %v", err)
		}

		if !resp.Success {
			t.Error("expected success = true")
		}

		// 验证商品被软删除
		var deletedProduct model.Product
		err = mysql.DB.Unscoped().First(&deletedProduct, testProduct.ID).Error
		if err != nil {
			t.Fatalf("failed to query deleted product: %v", err)
		}
		if deletedProduct.DeletedAt.Time.IsZero() {
			t.Error("product should be soft deleted")
		}

		// 验证关联的SKUs也被软删除
		var skus []model.SKU
		err = mysql.DB.Unscoped().Where("product_id = ?", testProduct.ID).Find(&skus).Error
		if err != nil {
			t.Fatalf("failed to query skus: %v", err)
		}
		if len(skus) != 2 {
			t.Errorf("got %d skus, want 2", len(skus))
		}
		for _, sku := range skus {
			if sku.DeletedAt.Time.IsZero() {
				t.Errorf("sku %d should be soft deleted", sku.ID)
			}
		}

		// 验证正常查询无法查到已删除的商品
		err = mysql.DB.First(&model.Product{}, testProduct.ID).Error
		if err != gorm.ErrRecordNotFound {
			t.Errorf("expected ErrRecordNotFound, got %v", err)
		}
	})

	t.Run("product not found", func(t *testing.T) {
		svc := NewDeleteProductService(getTestContext())

		req := &product.DeleteProductReq{
			Id: 99999, // 不存在的商品ID
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for non-existent product")
		}
	})

	t.Run("already deleted", func(t *testing.T) {
		// 先删除一次
		firstSvc := NewDeleteProductService(getTestContext())
		firstReq := &product.DeleteProductReq{Id: int64(testProduct.ID)}
		_, _ = firstSvc.Run(firstReq)

		// 尝试再次删除
		secondSvc := NewDeleteProductService(getTestContext())
		secondReq := &product.DeleteProductReq{Id: int64(testProduct.ID)}

		resp, err := secondSvc.Run(secondReq)
		if err == nil {
			t.Error("expected error for already deleted product, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for already deleted product")
		}
	})

	t.Run("verify related data integrity", func(t *testing.T) {
		// 创建新的测试商品和SKU
		newProduct := createTestProduct(t, int64(category.ID))
		_ = createTestSKU(t, newProduct.ID)

		// 删除商品
		svc := NewDeleteProductService(getTestContext())
		req := &product.DeleteProductReq{Id: int64(newProduct.ID)}
		_, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to delete product: %v", err)
		}

		// 验证关联数据的完整性
		var count int64
		err = mysql.DB.Model(&model.SKU{}).Where("product_id = ?", newProduct.ID).Count(&count).Error
		if err != nil {
			t.Fatalf("failed to count skus: %v", err)
		}
		if count > 0 {
			t.Error("expected no active SKUs for deleted product")
		}
	})
}
