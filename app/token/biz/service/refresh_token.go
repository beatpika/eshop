package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/beatpika/eshop/app/token/biz/dal/redis"
	"github.com/beatpika/eshop/app/token/biz/utils"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type RefreshTokenService struct {
	ctx context.Context
}

func NewRefreshTokenService(ctx context.Context) *RefreshTokenService {
	initJWTUtil() // 确保JWT工具已初始化
	return &RefreshTokenService{ctx: ctx}
}

func (s *RefreshTokenService) Run(req *auth.RefreshTokenRequest) (resp *auth.RefreshTokenResponse, err error) {
	resp = &auth.RefreshTokenResponse{
		ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
	}

	// 从Redis获取refresh token信息
	refreshTokenData, err := redis.TokenManager.GetRefreshToken(s.ctx, req.RefreshToken)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid refresh token: %v", err)
		return resp, nil
	}

	// 解析存储的数据
	userID, err := strconv.ParseInt(refreshTokenData["user_id"], 10, 32)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid user_id in refresh token data: %v", err)
		return resp, nil
	}

	// 获取旧的access token
	oldAccessToken := refreshTokenData["access_token"]

	// 先删除旧的访问令牌
	if oldAccessToken != "" {
		_ = redis.TokenManager.DeleteAccessToken(s.ctx, oldAccessToken, int32(userID))
		// 等待删除操作完成
		time.Sleep(50 * time.Millisecond)

		// 确认旧token已被删除
		_, err := redis.TokenManager.GetAccessToken(s.ctx, oldAccessToken)
		if err == nil {
			resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
			resp.ErrorMessage = "failed to delete old access token"
			return resp, nil
		}
	}

	// 获取用户角色（默认为普通用户）
	role := int64(auth.UserRole_USER_ROLE_USER)

	// 生成新的访问令牌
	accessToken, expiresAt, err := s.generateNewAccessToken(int32(userID), auth.UserRole(role))
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to generate new access token: %v", err)
		return resp, nil
	}

	// 保存新的访问令牌
	err = redis.TokenManager.SaveAccessToken(s.ctx, accessToken, int32(userID), int32(role), expiresAt)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to save new access token: %v", err)
		return resp, nil
	}

	// 更新refresh token关联的access token
	err = redis.TokenManager.SaveRefreshToken(s.ctx, req.RefreshToken, accessToken, int32(userID), expiresAt)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to update refresh token: %v", err)
		return resp, nil
	}

	// 填充响应
	resp.AccessToken = accessToken
	resp.ExpiresAt = expiresAt
	resp.RefreshToken = req.RefreshToken // 返回相同的refresh token
	resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
	resp.ErrorMessage = ""

	return resp, nil
}

// generateNewAccessToken 生成新的访问令牌
func (s *RefreshTokenService) generateNewAccessToken(userID int32, role auth.UserRole) (string, int64, error) {
	now := time.Now()
	expiresAt := now.Add(utils.AccessTokenExpiry).Unix()
	token, err := jwtUtil.GenerateToken(userID, int32(role), utils.AccessTokenExpiry)
	if err != nil {
		return "", 0, fmt.Errorf("failed to generate JWT token: %w", err)
	}
	return token, expiresAt, nil
}
