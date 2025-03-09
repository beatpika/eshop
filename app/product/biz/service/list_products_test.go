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

// setupListTestDB 设置测试数据库
func setupListTestDB(t *testing.T, rows *sqlmock.Rows, total int64, queryErr error) (*gorm.DB, sqlmock.Sqlmock) {
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

	if queryErr != nil {
		mock.ExpectQuery("SELECT \\* FROM `products`").WillReturnError(queryErr)
	} else {
		mock.ExpectQuery("SELECT count\\(\\*\\) FROM `products`").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(total))

		mock.ExpectQuery("SELECT \\* FROM `products`").WillReturnRows(rows)
	}

	return db, mock
}

func TestListProducts_Run(t *testing.T) {
	tests := []struct {
		name       string
		req        *product.ListProductsReq
		mockRows   *sqlmock.Rows
		mockTotal  int64
		queryErr   error
		wantErr    bool
		wantTotal  int32
		wantLength int
	}{
		{
			name: "success",
			req: &product.ListProductsReq{
				Page:     1,
				PageSize: 10,
				Category: "electronics",
			},
			mockRows: sqlmock.NewRows([]string{"id", "name", "description", "picture", "price", "categories"}).
				AddRow(1, "Product 1", "Desc 1", "pic1.jpg", 99.99, `["electronics"]`).
				AddRow(2, "Product 2", "Desc 2", "pic2.jpg", 199.99, `["electronics"]`),
			mockTotal:  2,
			queryErr:   nil,
			wantErr:    false,
			wantTotal:  2,
			wantLength: 2,
		},
		{
			name: "empty_list",
			req: &product.ListProductsReq{
				Page:     1,
				PageSize: 10,
				Category: "nonexistent",
			},
			mockRows:   sqlmock.NewRows([]string{"id", "name", "description", "picture", "price", "categories"}),
			mockTotal:  0,
			queryErr:   nil,
			wantErr:    false,
			wantTotal:  0,
			wantLength: 0,
		},
		{
			name: "invalid_json",
			req: &product.ListProductsReq{
				Page:     1,
				PageSize: 10,
			},
			mockRows: sqlmock.NewRows([]string{"id", "name", "description", "picture", "price", "categories"}).
				AddRow(1, "Product 1", "Desc 1", "pic1.jpg", 99.99, "{invalid json}"),
			mockTotal: 1,
			queryErr:  nil,
			wantErr:   true,
		},
		{
			name: "database_error",
			req: &product.ListProductsReq{
				Page:     1,
				PageSize: 10,
			},
			mockRows:  nil,
			mockTotal: 0,
			queryErr:  gorm.ErrInvalidDB,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDB, mock := setupListTestDB(t, tt.mockRows, tt.mockTotal, tt.queryErr)
			originalDB := mysql.DB
			mysql.DB = testDB
			defer func() {
				mysql.DB = originalDB
			}()

			ctx := context.Background()
			s := NewListProductsService(ctx)
			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.wantTotal, resp.Total)
			assert.Equal(t, tt.wantLength, len(resp.Products))

			if len(resp.Products) > 0 {
				// 验证第一个商品的字段
				product := resp.Products[0]
				assert.Equal(t, uint32(1), product.Id)
				assert.Equal(t, "Product 1", product.Name)
				assert.Equal(t, "Desc 1", product.Description)
				assert.Equal(t, "pic1.jpg", product.Picture)
				assert.Equal(t, float32(99.99), product.Price)
				assert.Equal(t, []string{"electronics"}, product.Categories)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("database expectations were not met: %v", err)
			}
		})
	}
}
