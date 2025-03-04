package service

import (
	"context"
	"errors"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.UserRegisterReq) (resp *user.UserRegisterResp, err error) {
	resp = new(user.UserRegisterResp)
	resp.Base = new(common.BaseResp)

	// 检查密码是否匹配
	if req.Password != req.PasswordConfirm {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INVALID_REQUEST)
		resp.Base.StatusMessage = "Passwords do not match"
		return resp, errors.New("passwords do not match")
	}

	// 调用 user RPC 服务进行注册
	registerResp, err := rpc.UserClient.Register(h.Context, &user_pb.RegisterReq{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 生成登录令牌
	tokenResp, err := rpc.TokenClient.GenerateToken(h.Context, &auth.GenerateTokenRequest{
		UserId: registerResp.UserId,
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
	resp.UserId = registerResp.UserId
	resp.Token = tokenResp.AccessToken

	return resp, nil
}
