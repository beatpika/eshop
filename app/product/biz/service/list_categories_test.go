package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestListCategories(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("success - list root categories", func(t *testing.T) {
		// 创建多个一级分类，使用不同的排序序号
		_ = createTestCategory(t, "Category 1", 0, 1, 3) // name, parentID, level, sortOrder
		_ = createTestCategory(t, "Category 2", 0, 1, 1)
		_ = createTestCategory(t, "Category 3", 0, 1, 2)

		svc := NewListCategoriesService(getTestContext())

		req := &product.ListCategoriesReq{
			ParentId: 0, // 获取一级分类
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list categories: %v", err)
		}

		// 验证返回的分类数量
		if len(resp.Categories) != 3 {
			t.Errorf("got %d categories, want 3", len(resp.Categories))
		}

		// 验证排序是否正确（按照sort_order升序）
		for i := 0; i < len(resp.Categories)-1; i++ {
			if resp.Categories[i].SortOrder > resp.Categories[i+1].SortOrder {
				t.Errorf("categories not properly sorted by sort_order")
			}
		}
	})

	t.Run("success - list sub categories", func(t *testing.T) {
		// 创建父分类
		parentCategory := createTestCategory(t, "Parent Category", 0, 1, 1)

		// 创建子分类
		_ = createTestCategory(t, "Sub Category 1", int64(parentCategory.ID), 2, 2)
		_ = createTestCategory(t, "Sub Category 2", int64(parentCategory.ID), 2, 1)

		svc := NewListCategoriesService(getTestContext())

		req := &product.ListCategoriesReq{
			ParentId: int64(parentCategory.ID),
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list sub-categories: %v", err)
		}

		// 验证返回的分类数量
		if len(resp.Categories) != 2 {
			t.Errorf("got %d sub-categories, want 2", len(resp.Categories))
		}

		// 验证所有返回的分类都是指定父分类的子分类
		for _, category := range resp.Categories {
			if category.ParentId != int64(parentCategory.ID) {
				t.Errorf("category parent_id = %d, want %d", category.ParentId, parentCategory.ID)
			}
		}
	})

	t.Run("success - empty list", func(t *testing.T) {
		svc := NewListCategoriesService(getTestContext())

		req := &product.ListCategoriesReq{
			ParentId: 99999, // 不存在的父分类ID
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list categories: %v", err)
		}

		// 验证返回空列表而不是错误
		if resp.Categories == nil {
			t.Error("expected empty slice, got nil")
		}
		if len(resp.Categories) != 0 {
			t.Errorf("got %d categories, want 0", len(resp.Categories))
		}
	})

	t.Run("deleted categories not included", func(t *testing.T) {
		// 创建一个分类并软删除
		category := createTestCategory(t, "Deleted Category", 0, 1, 1)
		if err := mysql.DB.Delete(&category).Error; err != nil {
			t.Fatalf("failed to delete category: %v", err)
		}

		svc := NewListCategoriesService(getTestContext())

		req := &product.ListCategoriesReq{
			ParentId: 0,
		}

		resp, err := svc.Run(req)
		if err != nil {
			t.Fatalf("failed to list categories: %v", err)
		}

		// 验证已删除的分类不在列表中
		for _, cat := range resp.Categories {
			if cat.Id == int64(category.ID) {
				t.Error("deleted category should not be included in list")
			}
		}
	})
}
