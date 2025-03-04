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

type DeactivateAccountService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeactivateAccountService(Context context.Context, RequestContext *app.RequestContext) *DeactivateAccountService {
	return &DeactivateAccountService{RequestContext: RequestContext, Context: Context}
}

func (h *DeactivateAccountService) Run(req *user.DeactivateAccountReq) (resp *user.DeactivateAccountResp, err error) {
	resp = new(user.DeactivateAccountResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务注销账号
	_, err = rpc.UserClient.DeactivateAccount(h.Context, &user_pb.DeactivateAccountReq{
		UserId:   req.UserId,
		Password: req.Password,
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 从请求头中获取 token 并撤销
	token := h.RequestContext.GetHeader("Authorization")
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
