package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *user.UserLogoutReq) (resp *user.UserLogoutResp, err error) {
	resp = new(user.UserLogoutResp)
	resp.Base = new(common.BaseResp)

	// 从请求头中获取 token
	token := h.RequestContext.GetHeader("Authorization")

	// 调用 token RPC 服务撤销 token
	_, err = rpc.TokenClient.RevokeToken(h.Context, &auth.RevokeTokenRequest{
		Token: string(token),
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 设置响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "Success"

	return resp, nil
}
