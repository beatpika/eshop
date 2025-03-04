package service

import (
	"testing"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestLoginService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	// 创建一个禁用状态的用户
	disabledUser := &model.User{
		Email:          "disabled@example.com",
		Username:       "disableduser",
		PasswordHashed: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // password123
		Status:         2,                                                              // 禁用状态
	}
	if err := mysql.DB.Create(disabledUser).Error; err != nil {
		t.Fatalf("failed to create disabled test user: %v", err)
	}

	tests := []struct {
		name    string
		req     *user.LoginReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "successful login",
			req: &user.LoginReq{
				Email:    testUser.Email,
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "wrong password",
			req: &user.LoginReq{
				Email:    testUser.Email,
				Password: "wrongpassword",
			},
			wantErr: true,
			errMsg:  "invalid email or password",
		},
		{
			name: "user not found",
			req: &user.LoginReq{
				Email:    "nonexistent@example.com",
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "invalid email or password",
		},
		{
			name: "disabled account",
			req: &user.LoginReq{
				Email:    disabledUser.Email,
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "user account is disabled",
		},
		{
			name: "empty email",
			req: &user.LoginReq{
				Email:    "",
				Password: "password123",
			},
			wantErr: true,
			errMsg:  "email or password is empty",
		},
		{
			name: "empty password",
			req: &user.LoginReq{
				Email:    testUser.Email,
				Password: "",
			},
			wantErr: true,
			errMsg:  "email or password is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewLoginService(ctx)
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

			// 验证返回的用户信息
			if resp.UserId != int32(testUser.ID) {
				t.Errorf("expected user ID %v, got %v", testUser.ID, resp.UserId)
			}

			if resp.Username != testUser.Username {
				t.Errorf("expected username %v, got %v", testUser.Username, resp.Username)
			}

			if resp.Email != testUser.Email {
				t.Errorf("expected email %v, got %v", testUser.Email, resp.Email)
			}
		})
	}
}
