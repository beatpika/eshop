package service

import (
	"context"
	"testing"
	"time"

	tokenRedis "github.com/beatpika/eshop/app/token/biz/dal/redis"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenService_Run(t *testing.T) {
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
	assert.NotEmpty(t, genResp.RefreshToken)

	oldAccessToken := genResp.AccessToken
	oldRefreshToken := genResp.RefreshToken

	svc := NewRefreshTokenService(ctx)

	t.Run("使用有效的refresh_token", func(t *testing.T) {
		// 验证初始token存在
		_, err = tokenRedis.TokenManager.GetAccessToken(ctx, oldAccessToken)
		assert.NoError(t, err, "初始access token应该存在")
		_, err = tokenRedis.TokenManager.GetRefreshToken(ctx, oldRefreshToken)
		assert.NoError(t, err, "初始refresh token应该存在")

		// 执行刷新
		resp, err := svc.Run(&auth.RefreshTokenRequest{
			RefreshToken: oldRefreshToken,
		})
		assert.NoError(t, err)

		// 基本响应检查
		assert.NotEmpty(t, resp.AccessToken, "新的access token不应为空")
		assert.NotEqual(t, oldAccessToken, resp.AccessToken, "新旧access token不应相同")
		assert.NotEmpty(t, resp.RefreshToken, "新的refresh token不应为空")
		assert.NotEqual(t, oldRefreshToken, resp.RefreshToken, "新旧refresh token不应相同")
		assert.Greater(t, resp.ExpiresAt, time.Now().Unix())
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_UNSPECIFIED, resp.ErrorCode)
		assert.Empty(t, resp.ErrorMessage)

		// 验证新token在Redis中的状态
		tokenData, err := tokenRedis.TokenManager.GetAccessToken(ctx, resp.AccessToken)
		assert.NoError(t, err)
		assert.NotNil(t, tokenData)
		assert.Equal(t, "123", tokenData["user_id"])

		// 验证新的refresh token在Redis中的状态
		refreshData, err := tokenRedis.TokenManager.GetRefreshToken(ctx, resp.RefreshToken)
		assert.NoError(t, err)
		assert.NotNil(t, refreshData)
		assert.Equal(t, "123", refreshData["user_id"])
		assert.Equal(t, resp.AccessToken, refreshData["access_token"])

		// 验证旧token已从Redis中删除
		_, err = tokenRedis.TokenManager.GetAccessToken(ctx, oldAccessToken)
		assert.Error(t, err, "旧的access token应该已被删除")
		_, err = tokenRedis.TokenManager.GetRefreshToken(ctx, oldRefreshToken)
		assert.Error(t, err, "旧的refresh token应该已被删除")

		// 尝试使用旧的refresh token（应该失败）
		oldResp, err := svc.Run(&auth.RefreshTokenRequest{
			RefreshToken: oldRefreshToken,
		})
		assert.NoError(t, err)
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_INVALID_TOKEN, oldResp.ErrorCode)
		assert.Empty(t, oldResp.AccessToken)
	})

	t.Run("使用无效的refresh_token", func(t *testing.T) {
		resp, err := svc.Run(&auth.RefreshTokenRequest{
			RefreshToken: "invalid-refresh-token",
		})
		assert.NoError(t, err)
		assert.Empty(t, resp.AccessToken)
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_INVALID_TOKEN, resp.ErrorCode)
	})

	t.Run("使用过期的refresh_token", func(t *testing.T) {
		// 使token在Redis中过期
		mr.FastForward(31 * 24 * time.Hour)

		resp, err := svc.Run(&auth.RefreshTokenRequest{
			RefreshToken: genResp.RefreshToken,
		})
		assert.NoError(t, err)
		assert.Empty(t, resp.AccessToken)
		assert.Equal(t, auth.ErrorCode_ERROR_CODE_INVALID_TOKEN, resp.ErrorCode)
	})
}

func TestRefreshTokenService_generateNewAccessToken(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	svc := NewRefreshTokenService(context.Background())

	userID := int32(123)
	role := auth.UserRole_USER_ROLE_USER

	token, expiresAt, err := svc.generateNewAccessToken(userID, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 验证过期时间在合理范围内
	now := time.Now().Unix()
	assert.Greater(t, expiresAt, now)
	assert.Less(t, expiresAt, now+(3*60*60)) // 不应超过3小时

	// 验证token可以被解码并包含正确的信息
	claims, err := jwtUtil.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, int32(role), claims.Role)
}
