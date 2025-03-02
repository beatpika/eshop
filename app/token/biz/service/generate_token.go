package service

import (
	"context"
	"fmt"
	"time"

	"github.com/beatpika/eshop/app/token/biz/dal/redis"
	"github.com/beatpika/eshop/app/token/biz/utils"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/google/uuid"
)

type GenerateTokenService struct {
	ctx context.Context
}

func NewGenerateTokenService(ctx context.Context) *GenerateTokenService {
	return &GenerateTokenService{
		ctx: ctx,
	}
}

func (s *GenerateTokenService) Run(req *auth.GenerateTokenRequest) (resp *auth.GenerateTokenResponse, err error) {
	resp = &auth.GenerateTokenResponse{
		ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
	}

	// 检查JWT工具是否已初始化
	if utils.DefaultJWTUtil == nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = "JWT utility not initialized"
		return resp, nil
	}

	// 生成访问令牌
	accessToken, expiresAt, err := s.generateAccessToken(req.UserId, req.Role)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to generate access token: %v", err)
		return resp, nil
	}

	// 生成刷新令牌
	refreshToken, refreshExpiresAt, err := s.generateRefreshToken()
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to generate refresh token: %v", err)
		return resp, nil
	}

	// 保存访问令牌到Redis
	err = redis.TokenManager.SaveAccessToken(s.ctx, accessToken, req.UserId, int32(req.Role), expiresAt)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to save access token: %v", err)
		return resp, nil
	}

	// 保存刷新令牌到Redis和关联的访问令牌
	err = redis.TokenManager.SaveRefreshToken(s.ctx, refreshToken, accessToken, req.UserId, refreshExpiresAt)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
		resp.ErrorMessage = fmt.Sprintf("failed to save refresh token: %v", err)
		return resp, nil
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	resp.ExpiresAt = expiresAt

	return resp, nil
}

// generateAccessToken 生成访问令牌
func (s *GenerateTokenService) generateAccessToken(userID int32, role auth.UserRole) (string, int64, error) {
	expiresAt := time.Now().Add(utils.AccessTokenExpiry).Unix()
	token, err := utils.DefaultJWTUtil.GenerateToken(userID, int32(role), utils.AccessTokenExpiry)
	if err != nil {
		return "", 0, fmt.Errorf("failed to generate JWT token: %w", err)
	}
	return token, expiresAt, nil
}

// generateRefreshToken 生成刷新令牌
func (s *GenerateTokenService) generateRefreshToken() (string, int64, error) {
	refreshToken := uuid.New().String()
	expiresAt := time.Now().Add(utils.RefreshTokenExpiry).Unix()
	return refreshToken, expiresAt, nil
}
