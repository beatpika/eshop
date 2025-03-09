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

// setupUpdateTestDB 设置测试数据库
func setupUpdateTestDB(t *testing.T, productExists bool, getErr, updateErr error) (*gorm.DB, sqlmock.Sqlmock) {
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

	if getErr != nil {
		mock.ExpectQuery("SELECT \\* FROM `products`").
			WillReturnError(getErr)
	} else if !productExists {
		mock.ExpectQuery("SELECT \\* FROM `products`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))
	} else {
		mock.ExpectQuery("SELECT \\* FROM `products`").
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "name", "description", "picture", "price", "categories",
			}).AddRow(
				1, "Old Name", "Old Description", "old.jpg", 99.99, `["old-category"]`,
			))

		if updateErr != nil {
			mock.ExpectExec("UPDATE `products`").WillReturnError(updateErr)
		} else {
			mock.ExpectExec("UPDATE `products`").WillReturnResult(sqlmock.NewResult(1, 1))
		}
	}

	return db, mock
}

func TestUpdateProduct_Run(t *testing.T) {
	tests := []struct {
		name          string
		req           *product.UpdateProductReq
		productExists bool
		getErr        error
		updateErr     error
		wantErr       bool
	}{
		{
			name: "success",
			req: &product.UpdateProductReq{
				Id:          1,
				Name:        "New Name",
				Description: "New Description",
				Picture:     "new.jpg",
				Price:       199.99,
				Categories:  []string{"new-category"},
			},
			productExists: true,
			getErr:        nil,
			updateErr:     nil,
			wantErr:       false,
		},
		{
			name: "product_not_found",
			req: &product.UpdateProductReq{
				Id: 999,
			},
			productExists: false,
			getErr:        gorm.ErrRecordNotFound,
			updateErr:     nil,
			wantErr:       true,
		},
		{
			name: "update_error",
			req: &product.UpdateProductReq{
				Id:          1,
				Name:        "New Name",
				Description: "New Description",
				Picture:     "new.jpg",
				Price:       199.99,
				Categories:  []string{"new-category"},
			},
			productExists: true,
			getErr:        nil,
			updateErr:     errors.New("database error"),
			wantErr:       true,
		},
		{
			name: "invalid_categories",
			req: &product.UpdateProductReq{
				Id:         1,
				Name:       "New Name",
				Categories: []string{"new-category", string([]byte{0xff, 0xfe})}, // 无效的UTF-8字符
			},
			productExists: true,
			getErr:        nil,
			updateErr:     nil,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB, mock := setupUpdateTestDB(t, tt.productExists, tt.getErr, tt.updateErr)
			originalDB := mysql.DB
			mysql.DB = testDB
			defer func() {
				mysql.DB = originalDB
			}()

			ctx := context.Background()
			s := NewUpdateProductService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Product)

			// 验证更新后的商品信息
			assert.Equal(t, tt.req.Id, resp.Product.Id)
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
