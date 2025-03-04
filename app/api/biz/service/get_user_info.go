package service

import (
	"context"

	"github.com/beatpika/eshop/app/api/hertz_gen/basic/user"
	"github.com/beatpika/eshop/app/api/hertz_gen/common"
	"github.com/beatpika/eshop/app/api/infra/rpc"
	user_pb "github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserInfoService {
	return &GetUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserInfoService) Run(req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp = new(user.GetUserInfoResp)
	resp.Base = new(common.BaseResp)

	// 调用 user RPC 服务获取用户信息
	userResp, err := rpc.UserClient.GetUserInfo(h.Context, &user_pb.GetUserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		resp.Base.StatusCode = int32(common.StatusCode_STATUS_INTERNAL_ERROR)
		resp.Base.StatusMessage = err.Error()
		return resp, err
	}

	// 转换用户信息
	resp.User = &user.UserInfo{
		UserId:    userResp.User.UserId,
		Username:  userResp.User.Username,
		Email:     userResp.User.Email,
		Phone:     userResp.User.Phone,
		Avatar:    userResp.User.Avatar,
		Address:   userResp.User.Address,
		Status:    userResp.User.Status,
		CreatedAt: userResp.User.CreatedAt,
	}

	// 设置响应
	resp.Base.StatusCode = int32(common.StatusCode_STATUS_OK)
	resp.Base.StatusMessage = "Success"

	return resp, nil
}
