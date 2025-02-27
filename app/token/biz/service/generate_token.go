package service

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
)

type GenerateTokenService struct {
	ctx context.Context
} // NewGenerateTokenService new GenerateTokenService
func NewGenerateTokenService(ctx context.Context) *GenerateTokenService {
	return &GenerateTokenService{ctx: ctx}
}

// Run create note info
func (s *GenerateTokenService) Run(req *auth.GenerateTokenRequest) (resp *auth.GenerateTokenResponse, err error) {
	// Finish your business logic.

	return
}
