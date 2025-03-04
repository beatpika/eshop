// 订单模型
package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email        string
	StreetAdress string
	City         string
	State        string
	Country      string
	Zipcode      int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100); uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11); not null;"`
	Consignee    Consignee   `gorm:"embedded"`                                   // embedded嵌入结构体
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"` // 关联订单商品表,外键为OrderIdRefer,指向Order的OrderId
	UserCurrency string      `gorm:"type:varchar(10); not null;"`
}

func (Order) TableName() string {
	return "order"
}

// 提供查询的方法
func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	if err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error; err != nil { // 预加载OrderItems,关联查询
		return nil, err
	}
	return orders, nil
}
