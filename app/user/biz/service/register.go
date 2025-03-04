package service

import (
	"context"
	"errors"
	"strings"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/app/user/biz/utils"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
)

var (
	ErrInvalidPassword = errors.New("password is too weak")
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidEmail    = errors.New("invalid email format")
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) Run(req *user.RegisterReq) (*user.RegisterResp, error) {
	// 验证输入
	if err := s.validateRegisterRequest(req); err != nil {
		return nil, err
	}

	// 检查邮箱是否已存在
	existingUser, err := model.GetByEmail(s.ctx, mysql.DB, req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailExists
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	newUser := &model.User{
		Email:          req.Email,
		Username:       req.Username,
		PasswordHashed: hashedPassword,
		Phone:          req.Phone,
		Status:         1, // 正常状态
	}

	if err := model.Create(s.ctx, mysql.DB, newUser); err != nil {
		return nil, err
	}

	// 返回响应
	return &user.RegisterResp{
		UserId:   int32(newUser.ID),
		Username: newUser.Username,
	}, nil
}

func (s *RegisterService) validateRegisterRequest(req *user.RegisterReq) error {
	// 验证邮箱
	if req.Email == "" || !isValidEmail(req.Email) {
		return ErrInvalidEmail
	}

	// 验证用户名
	if req.Username == "" {
		return errors.New("username is required")
	}

	// 验证密码强度
	if !utils.ValidatePassword(req.Password) {
		return ErrInvalidPassword
	}

	return nil
}

func isValidEmail(email string) bool {
	// 简单的邮箱格式验证
	// TODO: 使用更严格的邮箱验证规则
	if len(email) < 3 || !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}
	return true
}
