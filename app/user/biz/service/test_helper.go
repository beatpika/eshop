package service

import (
	"context"
	"testing"
	"time"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/app/user/biz/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func init() {
	var err error
	// 创建一个内存数据库
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: false,
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 初始化表结构
	if err := testDB.AutoMigrate(&model.User{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	// 设置全局DB
	mysql.DB = testDB
}

func setupTestDB(t *testing.T) func() {
	t.Helper()

	// 开始新的事务
	tx := testDB.Begin()
	if tx.Error != nil {
		t.Fatalf("failed to begin transaction: %v", tx.Error)
	}

	// 确保表存在
	if err := tx.AutoMigrate(&model.User{}); err != nil {
		tx.Rollback()
		t.Fatalf("failed to migrate database: %v", err)
	}

	// 保存原始DB并设置事务为当前DB
	originalDB := mysql.DB
	mysql.DB = tx

	return func() {
		defer func() {
			mysql.DB = originalDB
		}()

		// 回滚事务，确保数据清理
		if err := tx.Rollback().Error; err != nil {
			t.Errorf("failed to rollback transaction: %v", err)
		}

		// 清理主数据库中的数据
		if err := testDB.Session(&gorm.Session{AllowGlobalUpdate: true}).
			Delete(&model.User{}).Error; err != nil {
			t.Errorf("failed to clean test database: %v", err)
		}
	}
}

// createTestUser 创建测试用户
func createTestUser(t *testing.T) *model.User {
	t.Helper()

	// 生成密码哈希
	passwordHash, err := utils.HashPassword("password123")
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	user := &model.User{
		Email:          "test@example.com",
		Username:       "testuser",
		PasswordHashed: passwordHash,
		Phone:          "13800138000",
		Avatar:         "https://example.com/avatar.jpg",
		Address:        "Test Address",
		Status:         1,
	}

	if err := mysql.DB.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	// 验证用户已创建
	var createdUser model.User
	if err := mysql.DB.First(&createdUser, user.ID).Error; err != nil {
		t.Fatalf("failed to verify test user creation: %v", err)
	}

	// 验证密码哈希
	if !utils.CheckPasswordHash("password123", createdUser.PasswordHashed) {
		t.Fatal("password hash verification failed")
	}

	return &createdUser
}

// getTestContext 获取测试用的Context
func getTestContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
