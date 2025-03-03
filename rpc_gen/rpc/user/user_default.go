package user

import (
	"context"
	user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (resp *user.GetUserInfoResp, err error) {
	resp, err = defaultClient.GetUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = defaultClient.UpdateUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePassword(ctx context.Context, req *user.UpdatePasswordReq, callOptions ...callopt.Option) (resp *user.UpdatePasswordResp, err error) {
	resp, err = defaultClient.UpdatePassword(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePassword call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePhone(ctx context.Context, req *user.UpdatePhoneReq, callOptions ...callopt.Option) (resp *user.UpdatePhoneResp, err error) {
	resp, err = defaultClient.UpdatePhone(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePhone call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeactivateAccount(ctx context.Context, req *user.DeactivateAccountReq, callOptions ...callopt.Option) (resp *user.DeactivateAccountResp, err error) {
	resp, err = defaultClient.DeactivateAccount(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeactivateAccount call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
