package service

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/stretchr/testify/assert"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setupTestDB 设置测试数据库
func setupTestDB(t *testing.T, mockErr error) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	dialector := mysqldriver.New(mysqldriver.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	if mockErr != nil {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `products`").WillReturnError(mockErr)
		mock.ExpectRollback()
	} else {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `products`").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
	}

	return db, mock
}

func TestMain(m *testing.M) {
	mysql.DB = nil
}

func TestCreateProduct_Run(t *testing.T) {
	tests := []struct {
		name    string
		req     *product.CreateProductReq
		dbError error
		wantErr bool
	}{
		{
			name: "success",
			req: &product.CreateProductReq{
				Name:        "Test Product",
				Description: "Test Description",
				Picture:     "test.jpg",
				Price:       99.99,
				Categories:  []string{"electronics", "gadgets"},
			},
			dbError: nil,
			wantErr: false,
		},
		{
			name: "db_error",
			req: &product.CreateProductReq{
				Name:        "Test Product",
				Description: "Test Description",
				Picture:     "test.jpg",
				Price:       99.99,
				Categories:  []string{"electronics"},
			},
			dbError: errors.New("database error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB, mock := setupTestDB(t, tt.dbError)
			originalDB := mysql.DB
			mysql.DB = testDB
			defer func() {
				mysql.DB = originalDB
			}()

			ctx := context.Background()
			s := NewCreateProductService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Product)
			assert.Equal(t, tt.req.Name, resp.Product.Name)
			assert.Equal(t, tt.req.Description, resp.Product.Description)
			assert.Equal(t, tt.req.Picture, resp.Product.Picture)
			assert.Equal(t, tt.req.Price, resp.Product.Price)
			assert.Equal(t, tt.req.Categories, resp.Product.Categories)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("database expectations were not met: %v", err)
			}
		})
	}
}
