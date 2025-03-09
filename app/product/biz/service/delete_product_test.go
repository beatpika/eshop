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
func setupDeleteTestDB(t *testing.T, productExists bool, getErr, deleteErr error) (*gorm.DB, sqlmock.Sqlmock) {
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
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test Product"))
	}

	if deleteErr != nil {
		mock.ExpectExec("DELETE FROM `products`").
			WillReturnError(deleteErr)
	} else if productExists {
		mock.ExpectExec("DELETE FROM `products`").
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	return db, mock
}

func TestDeleteProduct_Run(t *testing.T) {
	tests := []struct {
		name          string
		req           *product.DeleteProductReq
		productExists bool
		getErr        error
		deleteErr     error
		wantErr       bool
	}{
		{
			name: "success",
			req: &product.DeleteProductReq{
				Id: 1,
			},
			productExists: true,
			getErr:        nil,
			deleteErr:     nil,
			wantErr:       false,
		},
		{
			name: "product_not_found",
			req: &product.DeleteProductReq{
				Id: 1,
			},
			productExists: false,
			getErr:        gorm.ErrRecordNotFound,
			deleteErr:     nil,
			wantErr:       true,
		},
		{
			name: "delete_error",
			req: &product.DeleteProductReq{
				Id: 1,
			},
			productExists: true,
			getErr:        nil,
			deleteErr:     errors.New("database error"),
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB, mock := setupDeleteTestDB(t, tt.productExists, tt.getErr, tt.deleteErr)
			originalDB := mysql.DB
			mysql.DB = testDB
			defer func() {
				mysql.DB = originalDB
			}()

			ctx := context.Background()
			s := NewDeleteProductService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.req.Id, resp.Id)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("database expectations were not met: %v", err)
			}
		})
	}
}
