package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdatePasswordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePasswordService(Context context.Context, RequestContext *app.RequestContext) *UpdatePasswordService {
	return &UpdatePasswordService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePasswordService) Run(req *user.UpdatePasswordReq) (resp *user.UpdatePasswordResp, err error) {
	resp = new(user.UpdatePasswordResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务更新密码
	_, err = rpc.UserClient.UpdatePassword(h.Context, &user_pb.UpdatePasswordReq{
		UserId:      req.UserId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
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
