package service

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type VerifyTokenService struct {
	ctx context.Context
} // NewVerifyTokenService new VerifyTokenService
func NewVerifyTokenService(ctx context.Context) *VerifyTokenService {
	return &VerifyTokenService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenService) Run(req *auth.VerifyTokenRequest) (resp *auth.VerifyTokenResponse, err error) {
	// Finish your business logic.

	return
}
