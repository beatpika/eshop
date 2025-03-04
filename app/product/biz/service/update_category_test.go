package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestUpdateCategory(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// 创建测试分类
	testCategory := createTestCategory(t, "Test Category", 0, 1, 1)

	t.Run("success", func(t *testing.T) {
		svc := NewUpdateCategoryService(getTestContext())

		req := &product.UpdateCategoryReq{
			Id:        int64(testCategory.ID),
			Name:      "Updated Category",
			SortOrder: 100,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to update category: %v", err)
		}

		// 验证响应
		if resp.Category == nil {
			t.Fatal("category in response is nil")
		}
		if resp.Category.Name != req.Name {
			t.Errorf("category name = %s, want %s", resp.Category.Name, req.Name)
		}
		if resp.Category.SortOrder != req.SortOrder {
			t.Errorf("category sort_order = %d, want %d", resp.Category.SortOrder, req.SortOrder)
		}

		// 验证分类层级和父分类ID保持不变
		if resp.Category.Level != testCategory.Level {
			t.Errorf("category level changed: got %d, want %d", resp.Category.Level, testCategory.Level)
		}
		if resp.Category.ParentId != int64(testCategory.ParentID) {
			t.Errorf("category parent_id changed: got %d, want %d", resp.Category.ParentId, testCategory.ParentID)
		}

		// 验证数据库记录
		var updatedCategory model.Category
		if err := mysql.DB.First(&updatedCategory, testCategory.ID).Error; err != nil {
			t.Fatalf("failed to query category: %v", err)
		}
		if updatedCategory.Name != req.Name {
			t.Errorf("db category name = %s, want %s", updatedCategory.Name, req.Name)
		}
		if updatedCategory.SortOrder != req.SortOrder {
			t.Errorf("db category sort_order = %d, want %d", updatedCategory.SortOrder, req.SortOrder)
		}
	})

	t.Run("category not found", func(t *testing.T) {
		svc := NewUpdateCategoryService(getTestContext())

		req := &product.UpdateCategoryReq{
			Id:        99999, // 不存在的分类ID
			Name:      "Updated Category",
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent category, got nil")
		}
		if resp != nil && resp.Category != nil {
			t.Error("expected nil category for non-existent category")
		}
	})

	t.Run("empty name", func(t *testing.T) {
		svc := NewUpdateCategoryService(getTestContext())

		req := &product.UpdateCategoryReq{
			Id:        int64(testCategory.ID),
			Name:      "", // 空名称
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for empty name, got nil")
		}
		if resp != nil && resp.Category != nil {
			t.Error("expected nil category for empty name")
		}

		// 验证数据库记录未被修改
		var category model.Category
		if err := mysql.DB.First(&category, testCategory.ID).Error; err != nil {
			t.Fatalf("failed to query category: %v", err)
		}
		if category.Name == "" {
			t.Error("category name should not be empty")
		}
	})

	t.Run("negative sort order", func(t *testing.T) {
		svc := NewUpdateCategoryService(getTestContext())

		req := &product.UpdateCategoryReq{
			Id:        int64(testCategory.ID),
			Name:      "Test Category",
			SortOrder: -1, // 负数排序序号
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for negative sort order, got nil")
		}
		if resp != nil && resp.Category != nil {
			t.Error("expected nil category for negative sort order")
		}

		// 验证数据库记录未被修改
		var category model.Category
		if err := mysql.DB.First(&category, testCategory.ID).Error; err != nil {
			t.Fatalf("failed to query category: %v", err)
		}
		if category.SortOrder < 0 {
			t.Error("category sort_order should not be negative")
		}
	})
}
