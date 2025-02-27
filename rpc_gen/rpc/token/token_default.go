package token

import (
	"context"
	auth "github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GenerateToken(ctx context.Context, req *auth.GenerateTokenRequest, callOptions ...callopt.Option) (resp *auth.GenerateTokenResponse, err error) {
	resp, err = defaultClient.GenerateToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GenerateToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, callOptions ...callopt.Option) (resp *auth.VerifyTokenResponse, err error) {
	resp, err = defaultClient.VerifyToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest, callOptions ...callopt.Option) (resp *auth.RefreshTokenResponse, err error) {
	resp, err = defaultClient.RefreshToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefreshToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest, callOptions ...callopt.Option) (resp *auth.RevokeTokenResponse, err error) {
	resp, err = defaultClient.RevokeToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RevokeToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
