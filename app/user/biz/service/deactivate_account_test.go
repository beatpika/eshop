package service

import (
	"testing"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestDeactivateAccountService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	tests := []struct {
		name    string
		req     *user.DeactivateAccountReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "wrong password",
			req: &user.DeactivateAccountReq{
				UserId:   int32(testUser.ID),
				Password: "wrongpassword",
			},
			wantErr: true,
			errMsg:  "invalid email or password",
		},
		{
			name: "successful account deactivation",
			req: &user.DeactivateAccountReq{
				UserId:   int32(testUser.ID),
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			req: &user.DeactivateAccountReq{
				UserId:   9999,
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewDeactivateAccountService(ctx)
			resp, err := s.Run(tt.req)

			// 检查错误
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error containing %v, got nil", tt.errMsg)
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("expected error message %v, got %v", tt.errMsg, err.Error())
				}
				return
			}

			// 检查成功情况
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if resp == nil {
				t.Error("expected non-nil response")
				return
			}

			if !resp.Success {
				t.Error("expected success response")
				return
			}

			// 验证用户状态
			deactivatedUser, err := model.GetByID(ctx, mysql.DB.Unscoped(), uint(tt.req.UserId))
			if err != nil {
				t.Errorf("failed to get deactivated user: %v", err)
				return
			}

			if deactivatedUser.Status != 2 {
				t.Errorf("expected user status to be 2 (disabled), got %v", deactivatedUser.Status)
			}

			if deactivatedUser.DeletedAt.Time.IsZero() {
				t.Error("expected user to be soft deleted")
			}
		})
	}
}

func TestDeactivateAccountService_TransactionRollback(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()
	testUser := createTestUser(t)

	// 模拟事务失败
	s := NewDeactivateAccountService(ctx)
	_, err := s.Run(&user.DeactivateAccountReq{
		UserId:   9999,
		Password: "password123",
	})

	if err == nil {
		t.Error("expected error, got nil")
		return
	}

	// 验证原始用户状态未被修改
	originalUser, err := model.GetByID(ctx, mysql.DB, uint(testUser.ID))
	if err != nil {
		t.Errorf("failed to get original user: %v", err)
		return
	}

	if originalUser.Status != 1 {
		t.Errorf("user status should not have changed, expected 1, got %v", originalUser.Status)
	}

	if !originalUser.DeletedAt.Time.IsZero() {
		t.Error("user should not have been soft deleted")
	}
}
