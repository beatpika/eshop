package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestDeleteSKU(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	// 创建测试SKU
	testSKU := createTestSKU(t, testProduct.ID)

	t.Run("success", func(t *testing.T) {
		// 先将库存设置为0
		if err := mysql.DB.Model(&testSKU).Update("stock", 0).Error; err != nil {
			t.Fatalf("failed to update stock: %v", err)
		}

		svc := NewDeleteSKUService(getTestContext())

		req := &product.DeleteSKUReq{
			Id: int64(testSKU.ID),
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to delete SKU: %v", err)
		}

		if !resp.Success {
			t.Error("expected success = true")
		}

		// 验证SKU已被软删除
		var deletedSKU model.SKU
		err = mysql.DB.Unscoped().First(&deletedSKU, testSKU.ID).Error
		if err != nil {
			t.Fatalf("failed to query deleted SKU: %v", err)
		}
		if deletedSKU.DeletedAt.Time.IsZero() {
			t.Error("SKU should be soft deleted")
		}

		// 验证正常查询无法查到已删除的SKU
		err = mysql.DB.First(&model.SKU{}, testSKU.ID).Error
		if err == nil {
			t.Error("expected error when querying deleted SKU")
		}
	})

	t.Run("SKU not found", func(t *testing.T) {
		svc := NewDeleteSKUService(getTestContext())

		req := &product.DeleteSKUReq{
			Id: 99999, // 不存在的SKU ID
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent SKU, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for non-existent SKU")
		}
	})

	t.Run("already deleted", func(t *testing.T) {
		// 创建并删除SKU
		sku := createTestSKU(t, testProduct.ID)
		// 将库存设置为0以允许删除
		if err := mysql.DB.Model(&sku).Update("stock", 0).Error; err != nil {
			t.Fatalf("failed to update stock: %v", err)
		}
		if err := mysql.DB.Delete(&sku).Error; err != nil {
			t.Fatalf("failed to soft delete SKU: %v", err)
		}

		svc := NewDeleteSKUService(getTestContext())

		req := &product.DeleteSKUReq{
			Id: int64(sku.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for already deleted SKU, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for already deleted SKU")
		}
	})

	t.Run("with stock remaining", func(t *testing.T) {
		// 创建有库存的SKU
		sku := createTestSKU(t, testProduct.ID)
		if err := mysql.DB.Model(&sku).Update("stock", 100).Error; err != nil {
			t.Fatalf("failed to update stock: %v", err)
		}

		svc := NewDeleteSKUService(getTestContext())

		req := &product.DeleteSKUReq{
			Id: int64(sku.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for SKU with remaining stock, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for SKU with remaining stock")
		}

		// 验证SKU未被删除
		var checkSKU model.SKU
		if err := mysql.DB.First(&checkSKU, sku.ID).Error; err != nil {
			t.Fatalf("SKU should still exist: %v", err)
		}
	})
}
