package main

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/beatpika/eshop/app/token/biz/service"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// GenerateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest) (resp *auth.GenerateTokenResponse, err error) {
	resp, err = service.NewGenerateTokenService(ctx).Run(req)

	return resp, err
}

// VerifyToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (resp *auth.VerifyTokenResponse, err error) {
	resp, err = service.NewVerifyTokenService(ctx).Run(req)

	return resp, err
}

// RefreshToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (resp *auth.RefreshTokenResponse, err error) {
	resp, err = service.NewRefreshTokenService(ctx).Run(req)

	return resp, err
}

// RevokeToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (resp *auth.RevokeTokenResponse, err error) {
	resp, err = service.NewRevokeTokenService(ctx).Run(req)

	return resp, err
}
