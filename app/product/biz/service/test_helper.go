package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	"github.com/beatpika/eshop/app/product/biz/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	testDB         *gorm.DB
	skuCodeCounter = 0
)

func init() {
	var err error
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	if err := testDB.AutoMigrate(&model.Product{}, &model.SKU{}, &model.Category{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	mysql.DB = testDB
}

func setupTestDB(t *testing.T) func() {
	t.Helper()

	tx := testDB.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	if err := tx.AutoMigrate(&model.Product{}, &model.SKU{}, &model.Category{}); err != nil {
		tx.Rollback()
		t.Fatalf("failed to migrate database: %v", err)
	}

	originalDB := mysql.DB
	mysql.DB = tx

	return func() {
		defer func() {
			mysql.DB = originalDB
		}()

		if err := tx.Rollback().Error; err != nil {
			t.Errorf("failed to rollback transaction: %v", err)
		}

		if err := testDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Product{}).Error; err != nil {
			t.Errorf("failed to clean products: %v", err)
		}
		if err := testDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.SKU{}).Error; err != nil {
			t.Errorf("failed to clean skus: %v", err)
		}
		if err := testDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Category{}).Error; err != nil {
			t.Errorf("failed to clean categories: %v", err)
		}
	}
}

// createTestCategory 创建测试分类
func createTestCategory(t *testing.T, name string, parentID int64, level int32, sortOrder int32) *model.Category {
	t.Helper()

	category := &model.Category{
		Name:      name,
		ParentID:  parentID,
		Level:     level,
		SortOrder: sortOrder,
	}

	if err := mysql.DB.Create(category).Error; err != nil {
		t.Fatalf("failed to create test category: %v", err)
	}

	return category
}

// createTestProduct 创建测试商品
func createTestProduct(t *testing.T, categoryID int64) *model.Product {
	t.Helper()

	images := []string{"http://example.com/image1.jpg", "http://example.com/image2.jpg"}
	imagesJSON, err := json.Marshal(images)
	if err != nil {
		t.Fatalf("failed to marshal images: %v", err)
	}

	product := &model.Product{
		Name:        "Test Product",
		Description: "Test Description",
		CategoryID:  categoryID,
		Images:      string(imagesJSON),
		Price:       1000, // 10.00
		Status:      1,    // 待上架
	}

	if err := mysql.DB.Create(product).Error; err != nil {
		t.Fatalf("failed to create test product: %v", err)
	}

	return product
}

// createTestSKU 创建测试SKU
func createTestSKU(t *testing.T, productID uint) *model.SKU {
	t.Helper()

	skuCodeCounter++
	specs := map[string]string{"color": "Red", "size": "L"}
	specsJSON, err := json.Marshal(specs)
	if err != nil {
		t.Fatalf("failed to marshal specs: %v", err)
	}

	sku := &model.SKU{
		ProductID: productID,
		Specs:     string(specsJSON),
		Price:     1000, // 10.00
		Stock:     100,
		Code:      fmt.Sprintf("TST-%03d", skuCodeCounter),
		Version:   1, // 初始化版本号
	}

	if err := mysql.DB.Create(sku).Error; err != nil {
		t.Fatalf("failed to create test sku: %v", err)
	}

	return sku
}

// getTestContext 获取测试用的Context
func getTestContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
