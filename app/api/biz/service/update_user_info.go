package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	resp = new(user.UpdateUserInfoResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务更新用户信息
	_, err = rpc.UserClient.UpdateUserInfo(h.Context, &user_pb.UpdateUserInfoReq{
		UserId:   req.UserId,
		Username: req.Username,
		Avatar:   req.Avatar,
		Address:  req.Address,
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
