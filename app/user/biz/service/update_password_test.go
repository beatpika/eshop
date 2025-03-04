package service

import (
	"testing"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/app/user/biz/utils"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestUpdatePasswordService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	tests := []struct {
		name    string
		req     *user.UpdatePasswordReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "incorrect old password",
			req: &user.UpdatePasswordReq{
				UserId:      int32(testUser.ID),
				OldPassword: "wrongpassword",
				NewPassword: "NewPassword123!",
			},
			wantErr: true,
			errMsg:  "old password is incorrect",
		},
		{
			name: "weak new password",
			req: &user.UpdatePasswordReq{
				UserId:      int32(testUser.ID),
				OldPassword: "password123",
				NewPassword: "123",
			},
			wantErr: true,
			errMsg:  "password is too weak",
		},
		{
			name: "same as old password",
			req: &user.UpdatePasswordReq{
				UserId:      int32(testUser.ID),
				OldPassword: "password123",
				NewPassword: "password123",
			},
			wantErr: true,
			errMsg:  "new password must be different from old password",
		},
		{
			name: "successful password update",
			req: &user.UpdatePasswordReq{
				UserId:      int32(testUser.ID),
				OldPassword: "password123",
				NewPassword: "NewPassword123!",
			},
			wantErr: false,
		},
		{
			name: "user not found",
			req: &user.UpdatePasswordReq{
				UserId:      9999,
				OldPassword: "password123",
				NewPassword: "NewPassword123!",
			},
			wantErr: true,
			errMsg:  "user not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewUpdatePasswordService(ctx)
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

			// 验证密码是否已更新
			updatedUser, err := model.GetByID(ctx, mysql.DB, uint(tt.req.UserId))
			if err != nil {
				t.Errorf("failed to get updated user: %v", err)
				return
			}

			// 验证新密码是否正确
			if !utils.CheckPasswordHash(tt.req.NewPassword, updatedUser.PasswordHashed) {
				t.Error("password was not updated correctly")
			}
		})
	}
}
