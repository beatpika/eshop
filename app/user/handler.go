package main

import (
	"context"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/beatpika/eshop/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp, err = service.NewGetUserInfoService(ctx).Run(req)

	return resp, err
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = service.NewUpdateUserInfoService(ctx).Run(req)

	return resp, err
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	resp, err = service.NewUpdatePasswordService(ctx).Run(req)

	return resp, err
}

// UpdatePhone implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePhone(ctx context.Context, req *user.UpdatePhoneReq) (resp *user.UpdatePhoneResp, err error) {
	resp, err = service.NewUpdatePhoneService(ctx).Run(req)

	return resp, err
}

// DeactivateAccount implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeactivateAccount(ctx context.Context, req *user.DeactivateAccountReq) (resp *user.DeactivateAccountResp, err error) {
	resp, err = service.NewDeactivateAccountService(ctx).Run(req)

	return resp, err
}
