package service

import (
	"testing"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestRegisterService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	tests := []struct {
		name    string
		req     *user.RegisterReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "successful registration",
			req: &user.RegisterReq{
				Email:    "new@example.com",
				Username: "newuser",
				Password: "Password123!",
				Phone:    "13800138001",
			},
			wantErr: false,
		},
		{
			name: "weak password",
			req: &user.RegisterReq{
				Email:    "test2@example.com",
				Username: "testuser2",
				Password: "123",
				Phone:    "13800138002",
			},
			wantErr: true,
			errMsg:  "password is too weak",
		},
		{
			name: "duplicate email",
			req: &user.RegisterReq{
				Email:    "test@example.com", // 已经在测试帮助函数中创建的用户
				Username: "testuser3",
				Password: "Password123!",
				Phone:    "13800138003",
			},
			wantErr: true,
			errMsg:  "email already exists",
		},
		{
			name: "invalid email format",
			req: &user.RegisterReq{
				Email:    "invalid-email",
				Username: "testuser4",
				Password: "Password123!",
				Phone:    "13800138004",
			},
			wantErr: true,
			errMsg:  "invalid email format",
		},
		{
			name: "empty username",
			req: &user.RegisterReq{
				Email:    "test5@example.com",
				Username: "",
				Password: "Password123!",
				Phone:    "13800138005",
			},
			wantErr: true,
			errMsg:  "username is required",
		},
	}

	// 先创建一个用户用于测试重复邮箱
	createTestUser(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRegisterService(ctx)
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

			if resp.UserId == 0 {
				t.Error("expected non-zero user ID")
			}

			if resp.Username != tt.req.Username {
				t.Errorf("expected username %v, got %v", tt.req.Username, resp.Username)
			}
		})
	}
}
