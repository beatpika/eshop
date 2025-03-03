package service

import (
	"context"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

type DeactivateAccountService struct {
	ctx context.Context
} // NewDeactivateAccountService new DeactivateAccountService
func NewDeactivateAccountService(ctx context.Context) *DeactivateAccountService {
	return &DeactivateAccountService{ctx: ctx}
}

// Run create note info
func (s *DeactivateAccountService) Run(req *user.DeactivateAccountReq) (resp *user.DeactivateAccountResp, err error) {
	// Finish your business logic.

	return
}
