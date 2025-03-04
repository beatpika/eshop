package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdatePhoneService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdatePhoneService(Context context.Context, RequestContext *app.RequestContext) *UpdatePhoneService {
	return &UpdatePhoneService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdatePhoneService) Run(req *user.UpdatePhoneReq) (resp *user.UpdatePhoneResp, err error) {
	resp = new(user.UpdatePhoneResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务更新手机号
	_, err = rpc.UserClient.UpdatePhone(h.Context, &user_pb.UpdatePhoneReq{
		UserId:     req.UserId,
		Phone:      req.Phone,
		VerifyCode: req.VerifyCode,
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
