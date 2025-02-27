package service

import (
	"context"
	"testing"
	"time"

	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTokenService_Run(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	mr, redisCleanup := setupTestRedis(t)
	defer redisCleanup()

	ctx := context.Background()
	userId := int32(123)
	role := auth.UserRole_USER_ROLE_USER

	// 先生成一个token用于测试
	genSvc := NewGenerateTokenService(ctx)
	genResp, err := genSvc.Run(&auth.GenerateTokenRequest{
		UserId: userId,
		Role:   role,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, genResp.AccessToken)

	tests := []struct {
		name  string
		token string
		want  struct {
			isValid      bool
			userId       int32
			role         auth.UserRole
			errorCode    auth.ErrorCode
			errorMessage string
		}
	}{
		{
			name:  "验证有效token",
			token: genResp.AccessToken,
			want: struct {
				isValid      bool
				userId       int32
				role         auth.UserRole
				errorCode    auth.ErrorCode
				errorMessage string
			}{
				isValid:      true,
				userId:       userId,
				role:         role,
				errorCode:    auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
				errorMessage: "",
			},
		},
		{
			name:  "验证无效token",
			token: "invalid-token",
			want: struct {
				isValid      bool
				userId       int32
				role         auth.UserRole
				errorCode    auth.ErrorCode
				errorMessage string
			}{
				isValid:   false,
				userId:    0,
				role:      auth.UserRole_USER_ROLE_UNSPECIFIED,
				errorCode: auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
			},
		},
	}

	svc := NewVerifyTokenService(ctx)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := svc.Run(&auth.VerifyTokenRequest{Token: tt.token})
			assert.NoError(t, err)

			assert.Equal(t, tt.want.isValid, resp.IsValid)
			if tt.want.isValid {
				assert.Equal(t, tt.want.userId, resp.UserId)
				assert.Equal(t, tt.want.role, resp.Role)
				assert.NotEmpty(t, resp.Permissions)
				assert.Greater(t, resp.ExpiresAt, time.Now().Unix())
			}
			assert.Equal(t, tt.want.errorCode, resp.ErrorCode)
			if tt.want.errorMessage != "" {
				assert.Equal(t, tt.want.errorMessage, resp.ErrorMessage)
			}
		})
	}

	// 测试过期token
	t.Run("验证过期token", func(t *testing.T) {
		// 让token在Redis中过期
		mr.FastForward(time.Hour * 3)

		// 验证token
		resp, err := svc.Run(&auth.VerifyTokenRequest{Token: genResp.AccessToken})
		assert.NoError(t, err)
		assert.False(t, resp.IsValid)
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_INVALID_TOKEN, resp.ErrorCode)
	})
}

func TestVerifyTokenService_getPermissionsByRole(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	svc := NewVerifyTokenService(context.Background())

	tests := []struct {
		name string
		role auth.UserRole
		want []string
	}{
		{
			name: "管理员权限",
			role: auth.UserRole_USER_ROLE_ADMIN,
			want: []string{
				"user:read", "user:write",
				"product:read", "product:write",
				"order:read", "order:write",
				"system:admin",
			},
		},
		{
			name: "商家权限",
			role: auth.UserRole_USER_ROLE_MERCHANT,
			want: []string{
				"user:read",
				"product:read", "product:write",
				"order:read",
			},
		},
		{
			name: "普通用户权限",
			role: auth.UserRole_USER_ROLE_USER,
			want: []string{
				"user:read",
				"product:read",
				"order:read",
			},
		},
		{
			name: "未指定角色权限",
			role: auth.UserRole_USER_ROLE_UNSPECIFIED,
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			perms := svc.getPermissionsByRole(tt.role)
			assert.ElementsMatch(t, tt.want, perms)
		})
	}
}
