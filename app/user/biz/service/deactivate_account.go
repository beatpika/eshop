package service

import (
	"context"
	"errors"
	"time"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/app/user/biz/utils"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

var ErrDeactivationFailed = errors.New("account deactivation failed")

type DeactivateAccountService struct {
	ctx context.Context
}

func NewDeactivateAccountService(ctx context.Context) *DeactivateAccountService {
	return &DeactivateAccountService{ctx: ctx}
}

func (s *DeactivateAccountService) Run(req *user.DeactivateAccountReq) (*user.DeactivateAccountResp, error) {
	var err error
	var userModel *model.User

	// 在事务外查询用户信息并验证密码
	userModel, err = model.GetByID(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 验证密码
	if !utils.CheckPasswordHash(req.Password, userModel.PasswordHashed) {
		return nil, ErrInvalidCredentials
	}

	// 开始事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 再次检查用户是否存在（使用事务）
		exists := true
		if err := tx.Model(&model.User{}).Where("id = ?", req.UserId).First(&model.User{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				exists = false
			} else {
				return err
			}
		}

		if !exists {
			return ErrUserNotFound
		}

		// 更新用户状态为禁用
		if err := tx.Model(&model.User{}).Where("id = ?", req.UserId).Update("status", 2).Error; err != nil {
			return err
		}

		// 手动设置软删除字段
		if err := tx.Model(&model.User{}).
			Where("id = ?", req.UserId).
			Update("deleted_at", time.Now()).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrDeactivationFailed
	}

	return &user.DeactivateAccountResp{
		Success: true,
	}, nil
}
