package service

import (
	"context"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

type UpdatePhoneService struct {
	ctx context.Context
} // NewUpdatePhoneService new UpdatePhoneService
func NewUpdatePhoneService(ctx context.Context) *UpdatePhoneService {
	return &UpdatePhoneService{ctx: ctx}
}

// Run create note info
func (s *UpdatePhoneService) Run(req *user.UpdatePhoneReq) (resp *user.UpdatePhoneResp, err error) {
	// Finish your business logic.

	return
}
