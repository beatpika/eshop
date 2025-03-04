package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestCreateCategory(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("success - create root category", func(t *testing.T) {
		svc := NewCreateCategoryService(getTestContext())

		req := &product.CreateCategoryReq{
			Name:      "Electronics",
			ParentId:  0, // 一级分类
			Level:     1,
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to create category: %v", err)
		}

		// 验证响应
		if resp.Category == nil {
			t.Fatal("category in response is nil")
		}
		if resp.Category.Name != req.Name {
			t.Errorf("category name = %s, want %s", resp.Category.Name, req.Name)
		}
		if resp.Category.Level != req.Level {
			t.Errorf("category level = %d, want %d", resp.Category.Level, req.Level)
		}

		// 验证数据库记录
		var category model.Category
		if err := mysql.DB.First(&category, resp.Category.Id).Error; err != nil {
			t.Fatalf("failed to query category: %v", err)
		}
		if category.Name != req.Name {
			t.Errorf("db category name = %s, want %s", category.Name, req.Name)
		}
	})

	t.Run("success - create sub category", func(t *testing.T) {
		// 创建父分类
		parentCategory := createTestCategory(t, "Electronics", 0, 1, 1)

		svc := NewCreateCategoryService(getTestContext())

		req := &product.CreateCategoryReq{
			Name:      "Phones",
			ParentId:  int64(parentCategory.ID),
			Level:     2,
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to create sub-category: %v", err)
		}

		// 验证响应
		if resp.Category == nil {
			t.Fatal("category in response is nil")
		}
		if resp.Category.ParentId != int64(parentCategory.ID) {
			t.Errorf("category parent_id = %d, want %d", resp.Category.ParentId, parentCategory.ID)
		}

		// 验证数据库记录
		var category model.Category
		if err := mysql.DB.First(&category, resp.Category.Id).Error; err != nil {
			t.Fatalf("failed to query category: %v", err)
		}
		if category.ParentID != int64(parentCategory.ID) {
			t.Errorf("db category parent_id = %d, want %d", category.ParentID, parentCategory.ID)
		}
	})

	t.Run("invalid level", func(t *testing.T) {
		svc := NewCreateCategoryService(getTestContext())

		invalidLevels := []int32{0, 4, -1}
		for _, level := range invalidLevels {
			req := &product.CreateCategoryReq{
				Name:      "Test Category",
				ParentId:  0,
				Level:     level,
				SortOrder: 1,
			}

			resp, err := svc.Run(req)
			if err == nil {
				t.Errorf("expected error for invalid level %d, got nil", level)
			}
			if resp != nil && resp.Category != nil {
				t.Errorf("expected nil category for invalid level %d", level)
			}
		}
	})

	t.Run("non-existent parent category", func(t *testing.T) {
		svc := NewCreateCategoryService(getTestContext())

		req := &product.CreateCategoryReq{
			Name:      "Test Category",
			ParentId:  99999, // 不存在的父分类ID
			Level:     2,
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent parent category, got nil")
		}
		if resp != nil && resp.Category != nil {
			t.Error("expected nil category for non-existent parent category")
		}
	})

	t.Run("invalid parent level", func(t *testing.T) {
		// 创建二级分类
		parentCategory := createTestCategory(t, "Level 2", 0, 2, 1)

		svc := NewCreateCategoryService(getTestContext())

		req := &product.CreateCategoryReq{
			Name:      "Test Category",
			ParentId:  int64(parentCategory.ID),
			Level:     2, // 尝试创建同级分类作为子分类，这是不允许的
			SortOrder: 1,
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for invalid parent level, got nil")
		}
		if resp != nil && resp.Category != nil {
			t.Error("expected nil category for invalid parent level")
		}
	})
}
