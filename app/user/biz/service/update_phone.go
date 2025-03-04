package service

import (
	"context"
	"errors"
	"regexp"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

var (
	ErrInvalidPhone      = errors.New("invalid phone number format")
	ErrInvalidVerifyCode = errors.New("invalid verification code")
	ErrPhoneExists       = errors.New("phone number already exists")
	phoneRegex           = regexp.MustCompile(`^1[3-9]\d{9}$`) // 中国大陆手机号格式
)

type UpdatePhoneService struct {
	ctx context.Context
}

func NewUpdatePhoneService(ctx context.Context) *UpdatePhoneService {
	return &UpdatePhoneService{ctx: ctx}
}

func (s *UpdatePhoneService) Run(req *user.UpdatePhoneReq) (*user.UpdatePhoneResp, error) {
	// 验证手机号格式
	if !isValidPhone(req.Phone) {
		return nil, ErrInvalidPhone
	}

	// 验证验证码
	if !s.verifyCode(req.Phone, req.VerifyCode) {
		return nil, ErrInvalidVerifyCode
	}

	// 在事务外验证用户是否存在
	if _, err := model.GetByID(s.ctx, mysql.DB, uint(req.UserId)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 在事务中执行更新
	err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 检查手机号是否已被其他用户使用
		var count int64
		if err := tx.Model(&model.User{}).
			Where("phone = ? AND id != ?", req.Phone, req.UserId).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return ErrPhoneExists
		}

		// 再次验证用户存在（使用事务）
		var user model.User
		if err := tx.Where("id = ?", req.UserId).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}

		// 更新手机号
		result := tx.Model(&model.User{}).
			Where("id = ?", req.UserId).
			Update("phone", req.Phone)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return ErrUserNotFound
		}

		return nil
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			return nil, ErrUserNotFound
		case errors.Is(err, ErrPhoneExists):
			return nil, ErrPhoneExists
		default:
			return nil, err
		}
	}

	return &user.UpdatePhoneResp{
		Success: true,
	}, nil
}

func isValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

func (s *UpdatePhoneService) verifyCode(phone, code string) bool {
	// TODO: 实现实际的短信验证码验证逻辑
	// 这里要求验证码长度为6位数字
	if len(code) != 6 {
		return false
	}
	for _, c := range code {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
