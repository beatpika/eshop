package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	Username       string `gorm:"type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
	Phone          string `gorm:"type:varchar(20)"`
	Avatar         string `gorm:"type:varchar(255)"`
	Address        string `gorm:"type:text"`
	Status         int    `gorm:"type:tinyint;default:1"` // 1:正常 2:禁用 3:待验证
}

func (User) TableName() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByID(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Save(user).Error
}

func UpdateUserStatus(ctx context.Context, db *gorm.DB, userID uint, status int) error {
	result := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdatePassword(ctx context.Context, db *gorm.DB, userID uint, passwordHashed string) error {
	result := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("password_hashed", passwordHashed)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdatePhone(ctx context.Context, db *gorm.DB, userID uint, phone string) error {
	result := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("phone", phone)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func SoftDeleteUser(ctx context.Context, db *gorm.DB, userID uint) error {
	result := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
