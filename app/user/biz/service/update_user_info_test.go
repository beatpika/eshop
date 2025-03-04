package service

import (
	"testing"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestUpdateUserInfoService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	tests := []struct {
		name    string
		req     *user.UpdateUserInfoReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "successful update all fields",
			req: &user.UpdateUserInfoReq{
				UserId:   int32(testUser.ID),
				Username: "newusername",
				Avatar:   "https://example.com/new-avatar.jpg",
				Address:  "New Test Address",
			},
			wantErr: false,
		},
		{
			name: "successful update partial fields",
			req: &user.UpdateUserInfoReq{
				UserId:   int32(testUser.ID),
				Username: "newusername2",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			req: &user.UpdateUserInfoReq{
				UserId:   9999,
				Username: "nonexistent",
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name: "username too short",
			req: &user.UpdateUserInfoReq{
				UserId:   int32(testUser.ID),
				Username: "a",
			},
			wantErr: true,
			errMsg:  "username cannot be empty",
		},
		{
			name: "invalid avatar URL",
			req: &user.UpdateUserInfoReq{
				UserId:   int32(testUser.ID),
				Username: "validname",
				Avatar:   "invalid-url",
			},
			wantErr: true,
			errMsg:  "invalid avatar URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewUpdateUserInfoService(ctx)
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

			// 验证更新是否生效
			updatedUser, err := model.GetByID(ctx, mysql.DB, uint(tt.req.UserId))
			if err != nil {
				t.Errorf("failed to get updated user: %v", err)
				return
			}

			if tt.req.Username != "" && updatedUser.Username != tt.req.Username {
				t.Errorf("username not updated, expected %v, got %v", tt.req.Username, updatedUser.Username)
			}

			if tt.req.Avatar != "" && updatedUser.Avatar != tt.req.Avatar {
				t.Errorf("avatar not updated, expected %v, got %v", tt.req.Avatar, updatedUser.Avatar)
			}

			if tt.req.Address != "" && updatedUser.Address != tt.req.Address {
				t.Errorf("address not updated, expected %v, got %v", tt.req.Address, updatedUser.Address)
			}
		})
	}
}
