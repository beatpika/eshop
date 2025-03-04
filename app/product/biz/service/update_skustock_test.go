package service

import (
	"sync"
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateSKUStock(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	category := createTestCategory(t, "Test Category", 0, 1, 1)

	// 创建测试商品
	testProduct := createTestProduct(t, int64(category.ID))

	// 创建测试SKU
	testSKU := createTestSKU(t, testProduct.ID)

	t.Run("success", func(t *testing.T) {
		svc := NewUpdateSKUStockService(getTestContext())

		req := &product.UpdateSKUStockReq{
			Id:    int64(testSKU.ID),
			Stock: 50,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update SKU stock: %v", err)
		}

		// 验证响应
		if resp.Sku == nil {
			t.Fatal("SKU in response is nil")
		}
		if resp.Sku.Stock != req.Stock {
			t.Errorf("SKU stock = %d, want %d", resp.Sku.Stock, req.Stock)
		}

		// 验证数据库记录
		var updatedSKU model.SKU
		if err := mysql.DB.First(&updatedSKU, testSKU.ID).Error; err != nil {
			t.Fatalf("failed to query updated SKU: %v", err)
		}
		if updatedSKU.Stock != req.Stock {
			t.Errorf("db SKU stock = %d, want %d", updatedSKU.Stock, req.Stock)
		}
		if updatedSKU.Version != testSKU.Version+1 {
			t.Errorf("db SKU version = %d, want %d", updatedSKU.Version, testSKU.Version+1)
		}
	})

	t.Run("SKU not found", func(t *testing.T) {
		svc := NewUpdateSKUStockService(getTestContext())

		req := &product.UpdateSKUStockReq{
			Id:    99999, // 不存在的SKU ID
			Stock: 50,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent SKU, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for non-existent SKU")
		}
	})

	t.Run("negative stock", func(t *testing.T) {
		svc := NewUpdateSKUStockService(getTestContext())

		req := &product.UpdateSKUStockReq{
			Id:    int64(testSKU.ID),
			Stock: -10,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for negative stock, got nil")
		}
		if resp != nil && resp.Sku != nil {
			t.Error("expected nil SKU for negative stock")
		}

		// 验证数据库记录未被修改
		var originalSKU model.SKU
		if err := mysql.DB.First(&originalSKU, testSKU.ID).Error; err != nil {
			t.Fatalf("failed to query original SKU: %v", err)
		}
		if originalSKU.Stock < 0 {
			t.Error("SKU stock should not be negative")
		}
	})

	t.Run("concurrent stock updates", func(t *testing.T) {
		// 创建新的SKU用于并发测试
		sku := createTestSKU(t, testProduct.ID)
		initialStock := int32(100)
		// 设置初始库存和版本
		mysql.DB.Model(&sku).Updates(map[string]interface{}{
			"stock":   initialStock,
			"version": 1,
		})

		var wg sync.WaitGroup
		concurrentUpdates := 10
		successCount := 0
		var mu sync.Mutex

		// 并发更新库存
		for i := 0; i < concurrentUpdates; i++ {
			wg.Add(1)
			go func(stockDelta int32) {
				defer wg.Done()

				svc := NewUpdateSKUStockService(getTestContext())
				req := &product.UpdateSKUStockReq{
					Id:    int64(sku.ID),
					Stock: initialStock + stockDelta,
				}

				if _, err := svc.Run(req); err == nil {
					mu.Lock()
					successCount++
					mu.Unlock()
				}
			}(int32(i + 1))
		}

		wg.Wait()

		// 验证最终库存状态
		var finalSKU model.SKU
		if err := mysql.DB.First(&finalSKU, sku.ID).Error; err != nil {
			t.Fatalf("failed to query final SKU: %v", err)
		}

		// 验证至少有一些更新成功，且版本号已更新
		if successCount == 0 {
			t.Error("expected some successful updates")
		}
		if finalSKU.Version <= 1 {
			t.Errorf("expected version > 1, got %d", finalSKU.Version)
		}
		if finalSKU.Stock <= initialStock {
			t.Errorf("expected stock > %d, got %d", initialStock, finalSKU.Stock)
		}
	})
}
