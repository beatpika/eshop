package service

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/beatpika/eshop/app/product/biz/dal/mysql"
	product "github.com/beatpika/eshop/rpc_gen/kitex_gen/product"
	"github.com/stretchr/testify/assert"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setupGetTestDB 设置测试数据库
func setupGetTestDB(t *testing.T, productExists bool, categories string) (*gorm.DB, sqlmock.Sqlmock) {
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

	if !productExists {
		mock.ExpectQuery("SELECT \\* FROM `products`").
			WillReturnError(gorm.ErrRecordNotFound)
	} else {
		rows := sqlmock.NewRows([]string{"id", "name", "description", "picture", "price", "categories"}).
			AddRow(1, "Test Product", "Test Description", "test.jpg", 99.99, categories)
		mock.ExpectQuery("SELECT \\* FROM `products`").
			WillReturnRows(rows)
	}

	return db, mock
}

func TestGetProduct_Run(t *testing.T) {
	tests := []struct {
		name          string
		req           *product.GetProductReq
		productExists bool
		categories    string
		wantErr       bool
	}{
		{
			name: "success",
			req: &product.GetProductReq{
				Id: 1,
			},
			productExists: true,
			categories:    `["electronics", "gadgets"]`,
			wantErr:       false,
		},
		{
			name: "product_not_found",
			req: &product.GetProductReq{
				Id: 999,
			},
			productExists: false,
			categories:    "",
			wantErr:       true,
		},
		{
			name: "invalid_categories_json",
			req: &product.GetProductReq{
				Id: 1,
			},
			productExists: true,
			categories:    "{invalid json}",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB, mock := setupGetTestDB(t, tt.productExists, tt.categories)
			originalDB := mysql.DB
			mysql.DB = testDB
			defer func() {
				mysql.DB = originalDB
			}()

			ctx := context.Background()
			s := NewGetProductService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Product)
			assert.Equal(t, tt.req.Id, resp.Product.Id)
			assert.Equal(t, "Test Product", resp.Product.Name)
			assert.Equal(t, "Test Description", resp.Product.Description)
			assert.Equal(t, "test.jpg", resp.Product.Picture)
			assert.Equal(t, float32(99.99), resp.Product.Price)
			assert.Equal(t, []string{"electronics", "gadgets"}, resp.Product.Categories)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("database expectations were not met: %v", err)
			}
		})
	}
}
