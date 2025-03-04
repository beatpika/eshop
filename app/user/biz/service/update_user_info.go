package service

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/beatpika/eshop/app/user/biz/dal/mysql"
	"github.com/beatpika/eshop/app/user/biz/model"
	"github.com/beatpika/eshop/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

var (
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrInvalidAvatar = errors.New("invalid avatar URL")
	ErrNoChange      = errors.New("no information to update")
)

type UpdateUserInfoService struct {
	ctx context.Context
}

func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (*user.UpdateUserInfoResp, error) {
	// 验证输入
	if err := s.validateUpdateRequest(req); err != nil {
		return nil, err
	}

	// 在事务外验证用户是否存在
	userModel, err := model.GetByID(s.ctx, mysql.DB, uint(req.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 在事务中执行更新
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 再次验证用户存在（使用事务）
		var user model.User
		if err := tx.Where("id = ?", req.UserId).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}

		// 构建更新字段
		updates := make(map[string]interface{})

		if req.Username != "" && req.Username != userModel.Username {
			updates["username"] = strings.TrimSpace(req.Username)
		}
		if req.Avatar != "" && req.Avatar != userModel.Avatar {
			updates["avatar"] = req.Avatar
		}
		if req.Address != "" && req.Address != userModel.Address {
			updates["address"] = strings.TrimSpace(req.Address)
		}

		// 如果没有需要更新的字段，直接返回成功
		if len(updates) == 0 {
			return nil
		}

		// 执行更新
		result := tx.Model(&model.User{}).
			Where("id = ?", req.UserId).
			Updates(updates)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return ErrUserNotFound
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user.UpdateUserInfoResp{
		Success: true,
	}, nil
}

func (s *UpdateUserInfoService) validateUpdateRequest(req *user.UpdateUserInfoReq) error {
	// 如果提供了用户名，验证格式
	if req.Username != "" {
		username := strings.TrimSpace(req.Username)
		if len(username) < 2 {
			return ErrEmptyUsername
		}
	}

	// 如果提供了头像URL，验证格式
	if req.Avatar != "" {
		if !isValidAvatarURL(req.Avatar) {
			return ErrInvalidAvatar
		}
	}

	// 如果提供了地址，验证格式
	if req.Address != "" {
		address := strings.TrimSpace(req.Address)
		if len(address) == 0 {
			req.Address = ""
		}
	}

	return nil
}

func isValidAvatarURL(avatarURL string) bool {
	if avatarURL == "" {
		return true
	}

	parsedURL, err := url.ParseRequestURI(avatarURL)
	if err != nil {
		return false
	}

	// 检查是否是HTTP(S)URL
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// 检查是否有主机名
	if parsedURL.Host == "" {
		return false
	}

	return true
}
