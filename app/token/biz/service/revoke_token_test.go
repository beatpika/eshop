package service

import (
	"context"
	"testing"
	"time"

	tokenRedis "github.com/beatpika/eshop/app/token/biz/dal/redis"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestRevokeTokenService_Run(t *testing.T) {
	mr, cleanup := setupTestRedis(t)
	defer cleanup()

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
			success      bool
			errorCode    auth.ErrorCode
			errorMessage string
		}
	}{
		{
			name:  "撤销有效token",
			token: genResp.AccessToken,
			want: struct {
				success      bool
				errorCode    auth.ErrorCode
				errorMessage string
			}{
				success:      true,
				errorCode:    auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
				errorMessage: "",
			},
		},
		{
			name:  "撤销无效token",
			token: "invalid-token",
			want: struct {
				success      bool
				errorCode    auth.ErrorCode
				errorMessage string
			}{
				success:   false,
				errorCode: auth.ErrorCode_ERROR_CODE_INVALID_TOKEN,
			},
		},
	}

	svc := NewRevokeTokenService(ctx)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := svc.Run(&auth.RevokeTokenRequest{
				Token: tt.token,
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want.success, resp.Success)
			assert.Equal(t, tt.want.errorCode, resp.ErrorCode)

			if tt.want.success {
				// 验证token已从Redis中删除
				_, err := tokenRedis.TokenManager.GetAccessToken(ctx, tt.token)
				assert.Error(t, err)

				// 验证相关的refresh token也应该被删除
				// 注意：当前实现中这个功能是TODO状态
			}
		})
	}

	// 测试撤销已过期token
	t.Run("撤销已过期token", func(t *testing.T) {
		mr.FastForward(3 * time.Hour) // 快进3小时使token过期

		resp, err := svc.Run(&auth.RevokeTokenRequest{
			Token: genResp.AccessToken,
		})
		assert.NoError(t, err)
		assert.False(t, resp.Success)
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_INVALID_TOKEN, resp.ErrorCode)
	})
}

func TestRevokeTokenService_findRefreshTokensByAccessToken(t *testing.T) {
	_, cleanup := setupTestRedis(t)
	defer cleanup()

	svc := NewRevokeTokenService(context.Background())

	// 测试未实现的功能
	tokens, err := svc.findRefreshTokensByAccessToken("test-token")
	assert.Error(t, err)
	assert.Nil(t, tokens)
	assert.Contains(t, err.Error(), "not implemented")
}
