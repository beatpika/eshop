package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateSKU(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	// 创建测试SKU
	testSKU := createTestSKU(t, testProduct.ID)

	t.Run("success", func(t *testing.T) {
		svc := NewUpdateSKUService(getTestContext())

		req := &product.UpdateSKUReq{
			Id: int64(testSKU.ID),
			Specs: map[string]string{
				"color": "Blue",
				"size":  "XL",
			},
			Price: 2000,
			Code:  "TST-002",
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update SKU: %v", err)
		}

		// 验证响应
		if resp.Sku == nil {
			t.Fatal("SKU in response is nil")
		}
		if resp.Sku.Price != req.Price {
			t.Errorf("SKU price = %d, want %d", resp.Sku.Price, req.Price)
		}
		if resp.Sku.Code != req.Code {
			t.Errorf("SKU code = %s, want %s", resp.Sku.Code, req.Code)
		}

		// 验证规格信息
		if len(resp.Sku.Specs) != len(req.Specs) {
			t.Errorf("got %d specs, want %d", len(resp.Sku.Specs), len(req.Specs))
		}
		for k, v := range req.Specs {
			if resp.Sku.Specs[k] != v {
				t.Errorf("spec[%s] = %s, want %s", k, resp.Sku.Specs[k], v)
			}
		}

		// 验证数据库记录
		var updatedSKU model.SKU
		if err := mysql.DB.First(&updatedSKU, testSKU.ID).Error; err != nil {
			t.Fatalf("failed to query updated SKU: %v", err)
		}
		if updatedSKU.Price != req.Price {
			t.Errorf("db SKU price = %d, want %d", updatedSKU.Price, req.Price)
		}
		if updatedSKU.Code != req.Code {
			t.Errorf("db SKU code = %s, want %s", updatedSKU.Code, req.Code)
		}
	})

	t.Run("SKU not found", func(t *testing.T) {
		svc := NewUpdateSKUService(getTestContext())

		req := &product.UpdateSKUReq{
			Id: 99999, // 不存在的SKU ID
			Specs: map[string]string{
				"color": "Blue",
			},
			Price: 1000,
			Code:  "TST-999",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent SKU, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for non-existent SKU")
		}
	})

	t.Run("duplicate SKU code", func(t *testing.T) {
		// 创建另一个SKU
		anotherSKU := createTestSKU(t, testProduct.ID)
		existingCode := anotherSKU.Code

		svc := NewUpdateSKUService(getTestContext())

		req := &product.UpdateSKUReq{
			Id: int64(testSKU.ID),
			Specs: map[string]string{
				"color": "Blue",
			},
			Price: 1000,
			Code:  existingCode, // 使用已存在的编码
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for duplicate SKU code, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for duplicate SKU code")
		}

		// 验证原SKU未被修改
		var originalSKU model.SKU
		if err := mysql.DB.First(&originalSKU, testSKU.ID).Error; err != nil {
			t.Fatalf("failed to query original SKU: %v", err)
		}
		if originalSKU.Code == existingCode {
			t.Error("SKU code should not be updated to duplicate value")
		}
	})

	t.Run("invalid specs format", func(t *testing.T) {
		svc := NewUpdateSKUService(getTestContext())

		req := &product.UpdateSKUReq{
			Id: int64(testSKU.ID),
			Specs: map[string]string{
				"": "Empty Key", // 无效的规格键
			},
			Price: 1000,
			Code:  "TST-003",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for invalid specs format, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for invalid specs format")
		}
	})

	t.Run("negative price", func(t *testing.T) {
		svc := NewUpdateSKUService(getTestContext())

		req := &product.UpdateSKUReq{
			Id:    int64(testSKU.ID),
			Specs: map[string]string{"color": "Red"},
			Price: -1000, // 负价格
			Code:  "TST-003",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for negative price, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for negative price")
		}
	})
}
