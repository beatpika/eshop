package service

import (
	"testing"

	"github.com/beatpika/eshop/app/product/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
)

func TestGetCategoryTree(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	t.Run("empty tree", func(t *testing.T) {
		svc := NewGetCategoryTreeService(getTestContext())
		resp, err := svc.Run(&product.GetCategoryTreeReq{})
		if err != nil {
			t.Fatalf("failed to get category tree: %v", err)
		}
		if len(resp.Categories) != 0 {
			t.Errorf("got %d categories, want 0", len(resp.Categories))
		}
	})

	t.Run("multi level categories", func(t *testing.T) {
		// 创建测试分类
		level1 := createTestCategory(t, "Electronics", 0, 1, 2)              // 一级分类
		level1_2 := createTestCategory(t, "Clothing", 0, 1, 1)               // 一级分类
		level2 := createTestCategory(t, "Phones", int64(level1.ID), 2, 1)    // 二级分类
		level2_2 := createTestCategory(t, "Laptops", int64(level1.ID), 2, 2) // 二级分类
		level3 := createTestCategory(t, "iPhones", int64(level2.ID), 3, 1)   // 三级分类
		level3_2 := createTestCategory(t, "Android", int64(level2.ID), 3, 2) // 三级分类

		svc := NewGetCategoryTreeService(getTestContext())
		resp, err := svc.Run(&product.GetCategoryTreeReq{})
		if err != nil {
			t.Fatalf("failed to get category tree: %v", err)
		}

		// 验证所有分类都被返回
		expectedCount := 6
		if len(resp.Categories) != expectedCount {
			t.Errorf("got %d categories, want %d", len(resp.Categories), expectedCount)
		}

		// 验证层级顺序
		for i := 0; i < len(resp.Categories)-1; i++ {
			curr := resp.Categories[i]
			next := resp.Categories[i+1]
			if curr.Level > next.Level {
				t.Errorf("incorrect level order at index %d: %d > %d", i, curr.Level, next.Level)
			}
		}

		// 验证排序顺序
		var lastSortOrder int32
		var lastLevel int32
		for _, cat := range resp.Categories {
			if cat.Level == lastLevel && cat.SortOrder < lastSortOrder {
				t.Errorf("incorrect sort order within level %d: %d < %d", cat.Level, cat.SortOrder, lastSortOrder)
			}
			lastSortOrder = cat.SortOrder
			lastLevel = cat.Level
		}

		// 验证层级关系
		validateCategory(t, resp.Categories, level1)
		validateCategory(t, resp.Categories, level1_2)
		validateCategory(t, resp.Categories, level2)
		validateCategory(t, resp.Categories, level2_2)
		validateCategory(t, resp.Categories, level3)
		validateCategory(t, resp.Categories, level3_2)
	})
}

// 辅助函数：验证分类信息
func validateCategory(t *testing.T, categories []*product.Category, expected *model.Category) {
	t.Helper()
	var found bool
	for _, cat := range categories {
		if cat.Id == int64(expected.ID) {
			found = true
			if cat.Name != expected.Name {
				t.Errorf("category name = %s, want %s", cat.Name, expected.Name)
			}
			if cat.ParentId != expected.ParentID {
				t.Errorf("category parent_id = %d, want %d", cat.ParentId, expected.ParentID)
			}
			if cat.Level != expected.Level {
				t.Errorf("category level = %d, want %d", cat.Level, expected.Level)
			}
			if cat.SortOrder != expected.SortOrder {
				t.Errorf("category sort_order = %d, want %d", cat.SortOrder, expected.SortOrder)
			}
			break
		}
	}
	if !found {
		t.Errorf("category with id %d not found", expected.ID)
	}
}
