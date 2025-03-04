package service

import (
	"testing"

	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestGetUserInfoService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	tests := []struct {
		name    string
		req     *user.GetUserInfoReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "successful get user info",
			req: &user.GetUserInfoReq{
				UserId: int32(testUser.ID),
			},
			wantErr: false,
		},
		{
			name: "user not found",
			req: &user.GetUserInfoReq{
				UserId: 9999,
			},
			wantErr: true,
			errMsg:  "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewGetUserInfoService(ctx)
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

			userInfo := resp.User
			if userInfo == nil {
				t.Error("expected non-nil user info")
				return
			}

			// 验证用户信息
			if userInfo.UserId != int32(testUser.ID) {
				t.Errorf("expected user ID %v, got %v", testUser.ID, userInfo.UserId)
			}

			if userInfo.Username != testUser.Username {
				t.Errorf("expected username %v, got %v", testUser.Username, userInfo.Username)
			}

			if userInfo.Email != testUser.Email {
				t.Errorf("expected email %v, got %v", testUser.Email, userInfo.Email)
			}

			if userInfo.Phone != testUser.Phone {
				t.Errorf("expected phone %v, got %v", testUser.Phone, userInfo.Phone)
			}

			if userInfo.Avatar != testUser.Avatar {
				t.Errorf("expected avatar %v, got %v", testUser.Avatar, userInfo.Avatar)
			}

			if userInfo.Address != testUser.Address {
				t.Errorf("expected address %v, got %v", testUser.Address, userInfo.Address)
			}

			if userInfo.Status != int32(testUser.Status) {
				t.Errorf("expected status %v, got %v", testUser.Status, userInfo.Status)
			}

			expectedCreatedAt := testUser.CreatedAt.Unix()
			if userInfo.CreatedAt != expectedCreatedAt {
				t.Errorf("expected created_at %v, got %v", expectedCreatedAt, userInfo.CreatedAt)
			}
		})
	}
}
