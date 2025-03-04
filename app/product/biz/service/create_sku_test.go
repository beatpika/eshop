package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestCreateSKU(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	t.Run("success", func(t *testing.T) {
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"color": "Blue",
				"size":  "XL",
			},
			Price: 1500,
			Stock: 50,
			Code:  "TST-NEW-001",
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to create SKU: %v", err)
		}

		// 验证响应
		if resp.Sku == nil {
			t.Fatal("SKU in response is nil")
		}
		if resp.Sku.ProductId != req.ProductId {
			t.Errorf("SKU product_id = %d, want %d", resp.Sku.ProductId, req.ProductId)
		}
		if resp.Sku.Price != req.Price {
			t.Errorf("SKU price = %d, want %d", resp.Sku.Price, req.Price)
		}
		if resp.Sku.Stock != req.Stock {
			t.Errorf("SKU stock = %d, want %d", resp.Sku.Stock, req.Stock)
		}
		if resp.Sku.Code != req.Code {
			t.Errorf("SKU code = %s, want %s", resp.Sku.Code, req.Code)
		}

		// 验证数据库记录
		var createdSKU model.SKU
		if err := mysql.DB.First(&createdSKU, resp.Sku.Id).Error; err != nil {
			t.Fatalf("failed to query created SKU: %v", err)
		}
		if createdSKU.ProductID != testProduct.ID {
			t.Errorf("db SKU product_id = %d, want %d", createdSKU.ProductID, testProduct.ID)
		}
		if createdSKU.Price != req.Price {
			t.Errorf("db SKU price = %d, want %d", createdSKU.Price, req.Price)
		}
		if createdSKU.Stock != req.Stock {
			t.Errorf("db SKU stock = %d, want %d", createdSKU.Stock, req.Stock)
		}
		if createdSKU.Code != req.Code {
			t.Errorf("db SKU code = %s, want %s", createdSKU.Code, req.Code)
		}
	})

	t.Run("duplicate SKU code", func(t *testing.T) {
		// 先创建一个SKU
		sku := createTestSKU(t, testProduct.ID)

		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"color": "Red",
				"size":  "M",
			},
			Price: 1000,
			Stock: 100,
			Code:  sku.Code, // 使用已存在的编码
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for duplicate SKU code, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for duplicate code")
		}
	})

	t.Run("product not found", func(t *testing.T) {
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: 99999, // 不存在的商品ID
			Specs: map[string]string{
				"color": "Red",
			},
			Price: 1000,
			Stock: 100,
			Code:  "TST-NEW-002",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent product, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for non-existent product")
		}
	})

	t.Run("invalid specs format", func(t *testing.T) {
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"": "Empty Key", // 无效的规格键
			},
			Price: 1000,
			Stock: 100,
			Code:  "TST-NEW-003",
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
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"color": "Red",
			},
			Price: -1000, // 负价格
			Stock: 100,
			Code:  "TST-NEW-004",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for negative price, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for negative price")
		}
	})

	t.Run("negative stock", func(t *testing.T) {
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"color": "Red",
			},
			Price: 1000,
			Stock: -10, // 负库存
			Code:  "TST-NEW-005",
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for negative stock, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for negative stock")
		}
	})

	t.Run("empty code", func(t *testing.T) {
		svc := NewCreateSKUService(getTestContext())

		req := &product.CreateSKUReq{
			ProductId: int64(testProduct.ID),
			Specs: map[string]string{
				"color": "Red",
			},
			Price: 1000,
			Stock: 100,
			Code:  "", // 空编码
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for empty code, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for empty code")
		}
	})
}
