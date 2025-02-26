package service

import (
	"context"
	"strings"

	user "github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	kitex_user "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
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
	// 创建默认响应
	resp = &user.UserLoginResp{
		Base: &common.BaseResp{
			StatusCode:    0, // 成功状态码
			StatusMessage: "success",
		},
	}

	// 基本参数验证
	if req.Email == "" || req.Password == "" {
		resp.Base.StatusCode = 400 // 无效请求
		resp.Base.StatusMessage = "邮箱和密码不能为空"
		return resp, nil
	}

	// 将 API 请求转换为 RPC 请求
	rpcReq := &kitex_user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}

	// 调用 RPC
	rpcResp, err := rpc.UserClient.Login(h.Context, rpcReq)
	if err != nil {
		// 处理不同类型的错误
		if strings.Contains(err.Error(), "crypto/bcrypt") {
			resp.Base.StatusCode = 401 // 未授权
			resp.Base.StatusMessage = "用户名或密码错误"
		} else {
			resp.Base.StatusCode = 500 // 内部错误
			resp.Base.StatusMessage = "登录服务暂时不可用，请稍后再试"
		}
		return resp, nil
	}

	// 成功情况
	resp.UserId = rpcResp.UserId
	resp.Token = "" // Token 字段可以根据需要设置

	return resp, nil
}
