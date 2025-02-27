package service

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type RevokeTokenService struct {
	ctx context.Context
} // NewRevokeTokenService new RevokeTokenService
func NewRevokeTokenService(ctx context.Context) *RevokeTokenService {
	return &RevokeTokenService{ctx: ctx}
}

// Run create note info
func (s *RevokeTokenService) Run(req *auth.RevokeTokenRequest) (resp *auth.RevokeTokenResponse, err error) {
	// Finish your business logic.

	return
}
