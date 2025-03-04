package model

import "gorm.io/gorm"

// 订单商品关联表
type OrderItem struct {
	gorm.Model
	ProductId    uint32  `gorm:"type:int(11) not null"`
	OrderIdRefer string  `gorm:"type:varchar(100);index"` // 和订单主表关联
	Quantity     uint32  `gorm:"type:int(11) not null"`
	Cost         float32 `gorm:"type:decimal(10,2) not null"`
}

// 表名的指定
func (OrderItem) TableName() string {
	return "order_item"
}
