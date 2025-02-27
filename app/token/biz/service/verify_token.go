package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beatpika/eshop/app/token/biz/dal/redis"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type VerifyTokenService struct {
	ctx context.Context
}

func NewVerifyTokenService(ctx context.Context) *VerifyTokenService {
	initJWTUtil() // 确保JWT工具已初始化
	return &VerifyTokenService{ctx: ctx}
}

func (s *VerifyTokenService) Run(req *auth.VerifyTokenRequest) (resp *auth.VerifyTokenResponse, err error) {
	resp = &auth.VerifyTokenResponse{
		IsValid:   false,
		ErrorCode: auth.ErrorCode_ERROR_CODE_UNSPECIFIED,
	}

	// 验证JWT token
	claims, err := jwtUtil.VerifyToken(req.Token)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid token: %v", err)
		return resp, nil
	}

	// 检查token是否过期
	if jwtUtil.IsTokenExpired(claims) {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_TOKEN_EXPIRED
		resp.ErrorMessage = "token expired"
		return resp, nil
	}

	// 从Redis获取token信息
	tokenData, err := redis.TokenManager.GetAccessToken(s.ctx, req.Token)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("token not found in store: %v", err)
		return resp, nil
	}

	// 解析Redis中存储的数据
	userID, err := strconv.ParseInt(tokenData["user_id"], 10, 32)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid user_id in token data: %v", err)
		return resp, nil
	}

	role, err := strconv.ParseInt(tokenData["role"], 10, 32)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid role in token data: %v", err)
		return resp, nil
	}

	expiresAt, err := strconv.ParseInt(tokenData["expires"], 10, 64)
	if err != nil {
		resp.ErrorCode = auth.ErrorCode_ERROR_CODE_INVALID_TOKEN
		resp.ErrorMessage = fmt.Sprintf("invalid expires_at in token data: %v", err)
		return resp, nil
	}

	// 填充响应
	resp.IsValid = true
	resp.UserId = int32(userID)
	resp.Role = auth.UserRole(role)
	resp.ExpiresAt = expiresAt
	resp.ErrorCode = auth.ErrorCode_ERROR_CODE_UNSPECIFIED
	resp.ErrorMessage = ""

	// 根据角色添加权限信息
	resp.Permissions = s.getPermissionsByRole(auth.UserRole(role))

	return resp, nil
}

// getPermissionsByRole 根据用户角色获取权限列表
func (s *VerifyTokenService) getPermissionsByRole(role auth.UserRole) []string {
	switch role {
	case auth.UserRole_USER_ROLE_ADMIN:
		return []string{
			"user:read", "user:write",
			"product:read", "product:write",
			"order:read", "order:write",
			"system:admin",
		}
	case auth.UserRole_USER_ROLE_MERCHANT:
		return []string{
			"user:read",
			"product:read", "product:write",
			"order:read",
		}
	case auth.UserRole_USER_ROLE_USER:
		return []string{
			"user:read",
			"product:read",
			"order:read",
		}
	default:
		return []string{}
	}
}
