package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	resp = new(user.UserLoginResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务进行登录
	loginResp, err := rpc.UserClient.Login(h.Context, &user_pb.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 生成登录令牌
	tokenResp, err := rpc.TokenClient.GenerateToken(h.Context, &auth.GenerateTokenRequest{
		UserId: loginResp.UserId,
		Role:   auth.UserRole_USER_ROLE_USER,
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 设置响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "Success"
	resp.UserId = loginResp.UserId
	resp.Token = tokenResp.AccessToken
	resp.Username = loginResp.Username
	resp.Avatar = loginResp.Avatar

	return resp, nil
}
