package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beatpika/eshop/app/token/biz/dal/redis"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type RevokeTokenService struct {
	ctx context.Context
}

func NewRevokeTokenService(ctx context.Context) *RevokeTokenService {
	initJWTUtil() // 确保JWT工具已初始化
	return &RevokeTokenService{ctx: ctx}
}

func (s *RevokeTokenService) Run(req *auth.RevokeTokenRequest) (resp *auth.RevokeTokenResponse, err error) {
	resp = &auth.RevokeTokenResponse{
		Success:   false,
		ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
	}

	// 从Redis获取token信息验证其有效性
	tokenData, err := redis.TokenManager.GetAccessToken(s.ctx, req.Token)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("token not found in store: %v", err)
		return resp, nil
	}

	// 验证JWT token
	_, err = jwtUtil.VerifyToken(req.Token)
	if err != nil {
		// 即使JWT验证失败，我们仍然会继续撤销token
		// 记录错误但不返回
	}

	// 获取用户ID
	userID, err := strconv.ParseInt(tokenData["user_id"], 10, 32)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid user_id in token data: %v", err)
		return resp, nil
	}

	// 删除访问令牌
	err = redis.TokenManager.DeleteAccessToken(s.ctx, req.Token, int32(userID))
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to revoke access token: %v", err)
		return resp, nil
	}

	// 查找并删除关联的refresh token
	refreshTokens, err := s.findRefreshTokensByAccessToken(req.Token)
	if err == nil {
		for _, refreshToken := range refreshTokens {
			_ = redis.TokenManager.DeleteRefreshToken(s.ctx, refreshToken)
		}
	}

	resp.Success = true
	return resp, nil
}

// findRefreshTokensByAccessToken 查找与access token关联的所有refresh tokens
func (s *RevokeTokenService) findRefreshTokensByAccessToken(accessToken string) ([]string, error) {
	// TODO: 实现在Redis中反向查找refresh tokens的逻辑
	// 这需要额外的数据结构来维护access token和refresh token的映射关系
	// 当前版本暂不实现此功能
	return nil, fmt.Errorf("not implemented")
}
