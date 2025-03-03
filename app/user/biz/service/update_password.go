package service

import (
	"context"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

type UpdatePasswordService struct {
	ctx context.Context
} // NewUpdatePasswordService new UpdatePasswordService
func NewUpdatePasswordService(ctx context.Context) *UpdatePasswordService {
	return &UpdatePasswordService{ctx: ctx}
}

// Run create note info
func (s *UpdatePasswordService) Run(req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	// Finish your business logic.

	return
}
