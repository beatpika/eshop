package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestDeleteCategory(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("success - delete leaf category", func(t *testing.T) {
		// 创建测试分类
		category := createTestCategory(t, "Test Category", 0, 1, 1)

		svc := NewDeleteCategoryService(getTestContext())

		req := &product.DeleteCategoryReq{
			Id: int64(category.ID),
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to delete category: %v", err)
		}

		if !resp.Success {
			t.Error("expected success = true")
		}

		// 验证分类已被软删除
		var deletedCategory model.Category
		err = mysql.DB.Unscoped().First(&deletedCategory, category.ID).Error
		if err != nil {
			t.Fatalf("failed to query deleted category: %v", err)
		}
		if deletedCategory.DeletedAt.Time.IsZero() {
			t.Error("category should be soft deleted")
		}

		// 验证正常查询无法查到已删除的分类
		err = mysql.DB.First(&model.Category{}, category.ID).Error
		if err == nil {
			t.Error("expected error when querying deleted category")
		}
	})

	t.Run("fail - category has sub-categories", func(t *testing.T) {
		// 创建父分类
		parentCategory := createTestCategory(t, "Parent Category", 0, 1, 1)
		// 创建子分类
		_ = createTestCategory(t, "Child Category", int64(parentCategory.ID), 2, 1)

		svc := NewDeleteCategoryService(getTestContext())

		req := &product.DeleteCategoryReq{
			Id: int64(parentCategory.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for category with sub-categories")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for category with sub-categories")
		}

		// 验证分类未被删除
		var category model.Category
		if err := mysql.DB.First(&category, parentCategory.ID).Error; err != nil {
			t.Fatalf("category should still exist: %v", err)
		}
	})

	t.Run("fail - category has products", func(t *testing.T) {
		// 创建分类
		category := createTestCategory(t, "Test Category", 0, 1, 1)
		// 创建关联的商品
		_ = createTestProduct(t, int64(category.ID))

		svc := NewDeleteCategoryService(getTestContext())

		req := &product.DeleteCategoryReq{
			Id: int64(category.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for category with products")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for category with products")
		}

		// 验证分类未被删除
		var cat model.Category
		if err := mysql.DB.First(&cat, category.ID).Error; err != nil {
			t.Fatalf("category should still exist: %v", err)
		}
	})

	t.Run("fail - category not found", func(t *testing.T) {
		svc := NewDeleteCategoryService(getTestContext())

		req := &product.DeleteCategoryReq{
			Id: 99999, // 不存在的分类ID
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for non-existent category")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for non-existent category")
		}
	})

	t.Run("fail - already deleted category", func(t *testing.T) {
		// 创建并删除分类
		category := createTestCategory(t, "Test Category", 0, 1, 1)
		if err := mysql.DB.Delete(&category).Error; err != nil {
			t.Fatalf("failed to soft delete category: %v", err)
		}

		svc := NewDeleteCategoryService(getTestContext())

		req := &product.DeleteCategoryReq{
			Id: int64(category.ID),
		}

		resp, err := svc.Run(req)
		if err == nil {
			t.Error("expected error for already deleted category")
		}
		if resp != nil && resp.Success {
			t.Error("expected success = false for already deleted category")
		}
	})
}
