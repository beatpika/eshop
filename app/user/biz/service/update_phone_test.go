package service

import (
	"testing"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

func TestUpdatePhoneService_Run(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	ctx := getTestContext()

	// 创建测试用户
	testUser := createTestUser(t)

	tests := []struct {
		name    string
		req     *user.UpdatePhoneReq
		wantErr bool
		errMsg  string
	}{
		{
			name: "successful phone update",
			req: &user.UpdatePhoneReq{
				UserId:     int32(testUser.ID),
				Phone:      "13900139000",
				VerifyCode: "123456",
			},
			wantErr: false,
		},
		{
			name: "invalid phone format",
			req: &user.UpdatePhoneReq{
				UserId:     int32(testUser.ID),
				Phone:      "invalid-phone",
				VerifyCode: "123456",
			},
			wantErr: true,
			errMsg:  "invalid phone number format",
		},
		{
			name: "invalid verify code",
			req: &user.UpdatePhoneReq{
				UserId:     int32(testUser.ID),
				Phone:      "13900139000",
				VerifyCode: "",
			},
			wantErr: true,
			errMsg:  "invalid verification code",
		},
		{
			name: "user not found",
			req: &user.UpdatePhoneReq{
				UserId:     9999,
				Phone:      "13900139000",
				VerifyCode: "123456",
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name: "phone number too short",
			req: &user.UpdatePhoneReq{
				UserId:     int32(testUser.ID),
				Phone:      "139",
				VerifyCode: "123456",
			},
			wantErr: true,
			errMsg:  "invalid phone number format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewUpdatePhoneService(ctx)
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

			// 验证手机号是否已更新
			updatedUser, err := model.GetByID(ctx, mysql.DB, uint(tt.req.UserId))
			if err != nil {
				t.Errorf("failed to get updated user: %v", err)
				return
			}

			if updatedUser.Phone != tt.req.Phone {
				t.Errorf("phone number not updated, expected %v, got %v", tt.req.Phone, updatedUser.Phone)
			}
		})
	}
}
