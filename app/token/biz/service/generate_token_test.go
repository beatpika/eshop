package service

import (
	"context"
	"testing"
	"time"

	tokenRedis "github.com/beatpika/eshop/app/token/biz/dal/redis"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenService_Run(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	mr, redisCleanup := setupTestRedis(t)
	defer redisCleanup()

	tests := []struct {
		name string
		req  *auth.GenerateTokenRequest
		want struct {
			hasError     bool
			errorCode    auth.ErrorCode
			errorMessage string
		}
	}{
		{
			name: "正常生成token",
			req: &auth.GenerateTokenRequest{
				UserId: 123,
				Role:   auth.UserRole_USER_ROLE_USER,
			},
			want: struct {
				hasError     bool
				errorCode    auth.ErrorCode
				errorMessage string
			}{
				hasError:     false,
				errorCode:    auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
				errorMessage: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			svc := NewGenerateTokenService(ctx)

			resp, err := svc.Run(tt.req)
			assert.Nil(t, err)

			// 验证基本响应
			assert.NotEmpty(t, resp.AccessToken)
			assert.NotEmpty(t, resp.RefreshToken)
			assert.Greater(t, resp.ExpiresAt, time.Now().Unix())
			assert.Equal(t, tt.want.errorCode, resp.ErrorCode)
			assert.Equal(t, tt.want.errorMessage, resp.ErrorMessage)

			// 验证Redis中的数据
			tokenData, err := tokenRedis.TokenManager.GetAccessToken(ctx, resp.AccessToken)
			assert.Nil(t, err)
			assert.NotNil(t, tokenData)

			refreshTokenData, err := tokenRedis.TokenManager.GetRefreshToken(ctx, resp.RefreshToken)
			assert.Nil(t, err)
			assert.NotNil(t, refreshTokenData)

			// 验证Redis中存储的用户ID
			assert.Equal(t, "123", tokenData["user_id"])
			assert.Equal(t, "123", refreshTokenData["user_id"])

			// 验证Redis中存储的访问令牌关联
			assert.Equal(t, resp.AccessToken, refreshTokenData["access_token"])

			// 清理测试数据
			mr.FlushAll()
		})
	}
}

func TestGenerateTokenService_generateAccessToken(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	svc := NewGenerateTokenService(context.Background())

	userID := int32(123)
	role := auth.UserRole_USER_ROLE_USER

	token, expiresAt, err := svc.generateAccessToken(userID, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 验证过期时间在合理范围内
	now := time.Now().Unix()
	assert.Greater(t, expiresAt, now)
	assert.Less(t, expiresAt, now+(3*60*60)) // 不应超过3小时

	// 验证token可以被解码
	claims, err := jwtUtil.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, int32(role), claims.Role)
}

func TestGenerateTokenService_generateRefreshToken(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	svc := NewGenerateTokenService(context.Background())

	refreshToken, expiresAt, err := svc.generateRefreshToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshToken)

	// 验证过期时间在合理范围内
	now := time.Now().Unix()
	assert.Greater(t, expiresAt, now)
	assert.Less(t, expiresAt, now+(31*24*60*60)) // 不应超过31天
}
