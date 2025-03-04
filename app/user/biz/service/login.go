package service

import (
	"context"
	"errors"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/app/user/biz/utils"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserDisabled       = errors.New("user account is disabled")
	ErrUserNotVerified    = errors.New("user account is not verified")
	ErrEmptyCredentials   = errors.New("email or password is empty")
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

func (s *LoginService) Run(req *user.LoginReq) (*user.LoginResp, error) {
	// 验证请求参数
	if req.Email == "" || req.Password == "" {
		return nil, ErrEmptyCredentials
	}

	// 根据邮箱查找用户
	userModel, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// 检查用户状态
	if userModel.Status == 2 { // 禁用状态
		return nil, ErrUserDisabled
	}
	if userModel.Status == 3 { // 待验证状态
		return nil, ErrUserNotVerified
	}
	if userModel.Status != 1 { // 非正常状态
		return nil, errors.New("invalid user status")
	}

	// 验证密码
	if !utils.CheckPasswordHash(req.Password, userModel.PasswordHashed) {
		return nil, ErrInvalidCredentials
	}

	// 返回用户信息
	return &user.LoginResp{
		UserId:   int32(userModel.ID),
		Username: userModel.Username,
		Email:    userModel.Email,
		Avatar:   userModel.Avatar,
	}, nil
}
