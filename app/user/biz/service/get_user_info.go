package service

import (
	"context"
	"errors"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

var ErrUserNotFound = errors.New("user not found")

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// 获取用户信息
	userModel, err := model.GetByID(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		return nil, ErrUserNotFound
	}

	// 构造用户信息响应
	userInfo := &user.UserInfo{
		UserId:    int32(userModel.ID),
		Username:  userModel.Username,
		Email:     userModel.Email,
		Phone:     userModel.Phone,
		Avatar:    userModel.Avatar,
		Address:   userModel.Address,
		Status:    int32(userModel.Status),
		CreatedAt: userModel.CreatedAt.Unix(),
	}

	return &user.GetUserInfoResp{
		User: userInfo,
	}, nil
}
