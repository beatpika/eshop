package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateProductStatus(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	t.Run("success - update to published", func(t *testing.T) {
		svc := NewUpdateProductStatusService(getTestContext())

		req := &product.UpdateProductStatusReq{
			Id:     int64(testProduct.ID),
			Status: 2, // 已上架
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update product status: %v", err)
		}

		if !resp.Success {
			t.Error("expected success = true")
		}

		// 验证数据库记录
		var updatedProduct model.Product
		if err := mysql.DB.First(&updatedProduct, testProduct.ID).Error; err != nil {
			t.Fatalf("failed to query updated product: %v", err)
		}
		if updatedProduct.Status != 2 {
			t.Errorf("product status = %d, want 2", updatedProduct.Status)
		}
	})

	t.Run("success - update to unpublished", func(t *testing.T) {
		svc := NewUpdateProductStatusService(getTestContext())

		req := &product.UpdateProductStatusReq{
			Id:     int64(testProduct.ID),
			Status: 3, // 已下架
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update product status: %v", err)
		}

		if !resp.Success {
			t.Error("expected success = true")
		}

		// 验证数据库记录
		var updatedProduct model.Product
		if err := mysql.DB.First(&updatedProduct, testProduct.ID).Error; err != nil {
			t.Fatalf("failed to query updated product: %v", err)
		}
		if updatedProduct.Status != 3 {
			t.Errorf("product status = %d, want 3", updatedProduct.Status)
		}
	})

	t.Run("product not found", func(t *testing.T) {
		svc := NewUpdateProductStatusService(getTestContext())

		req := &product.UpdateProductStatusReq{
			Id:     99999, // 不存在的商品ID
			Status: 2,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for non-existent product")
		}
	})

	t.Run("invalid status value", func(t *testing.T) {
		svc := NewUpdateProductStatusService(getTestContext())

		invalidStatuses := []int32{0, 4, -1, 100}

		for _, status := range invalidStatuses {
			req := &product.UpdateProductStatusReq{
				Id:     int64(testProduct.ID),
				Status: status,
			}

			resp, err := svc.Run(req)
			if err == nil {
				t.Errorf("expected error for invalid status %d, got nil", status)
			}
			if resp != nil && resp.Success {
				t.Errorf("expected success = false for invalid status %d", status)
			}

			// 验证数据库记录未被修改
			var product model.Product
			if err := mysql.DB.First(&product, testProduct.ID).Error; err != nil {
				t.Fatalf("failed to query product: %v", err)
			}
			if product.Status == status {
				t.Errorf("product status should not be updated to invalid value %d", status)
			}
		}
	})
}
