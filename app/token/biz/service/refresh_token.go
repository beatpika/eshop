package service

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type RefreshTokenService struct {
	ctx context.Context
} // NewRefreshTokenService new RefreshTokenService
func NewRefreshTokenService(ctx context.Context) *RefreshTokenService {
	return &RefreshTokenService{ctx: ctx}
}

// Run create note info
func (s *RefreshTokenService) Run(req *auth.RefreshTokenRequest) (resp *auth.RefreshTokenResponse, err error) {
	// Finish your business logic.

	return
}
