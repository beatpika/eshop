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
	ErrOldPasswordIncorrect = errors.New("old password is incorrect")
	ErrSamePassword         = errors.New("new password must be different from old password")
)

type UpdatePasswordService struct {
	ctx context.Context
}

func NewUpdatePasswordService(ctx context.Context) *UpdatePasswordService {
	return &UpdatePasswordService{ctx: ctx}
}

func (s *UpdatePasswordService) Run(req *user.UpdatePasswordReq) (*user.UpdatePasswordResp, error) {
	var err error
	var userModel *model.User

	// 在事务外查询用户信息
	userModel, err = model.GetByID(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 验证旧密码
	if !utils.CheckPasswordHash(req.OldPassword, userModel.PasswordHashed) {
		return nil, ErrOldPasswordIncorrect
	}

	// 验证新密码
	if err := s.validateNewPassword(req.NewPassword, req.OldPassword); err != nil {
		return nil, err
	}

	// 生成新密码的哈希
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}

	// 在事务中执行更新
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 再次验证用户存在和密码正确（使用事务）
		var user model.User
		if err := tx.Where("id = ?", req.UserId).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}

		if !utils.CheckPasswordHash(req.OldPassword, user.PasswordHashed) {
			return ErrOldPasswordIncorrect
		}

		// 更新密码
		result := tx.Model(&model.User{}).
			Where("id = ?", req.UserId).
			Update("password_hashed", hashedPassword)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return ErrUserNotFound
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, ErrOldPasswordIncorrect) {
			return nil, ErrOldPasswordIncorrect
		}
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user.UpdatePasswordResp{
		Success: true,
	}, nil
}

func (s *UpdatePasswordService) validateNewPassword(newPassword, oldPassword string) error {
	// 检查新密码是否与旧密码相同
	if newPassword == oldPassword {
		return ErrSamePassword
	}

	// 检查新密码强度
	if !utils.ValidatePassword(newPassword) {
		return ErrInvalidPassword
	}

	return nil
}
