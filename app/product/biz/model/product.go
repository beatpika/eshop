package model

import (
	"time"

	"gorm.io/gorm"
)

// Product 商品表
type Product struct {
	gorm.Model
	Name        string    // 商品名称
	Description string    // 商品描述
	CategoryID  int64     // 分类ID
	Images      string    // 商品图片，JSON数组
	Price       int64     // 商品价格（基准价格）
	Status      int32     // 商品状态：1-待上架，2-已上架，3-已下架
	SKUs        []SKU     // SKU列表
	Category    *Category // 商品分类
	Sales       int64     // 销量
}

// SKU 库存单位表
type SKU struct {
	gorm.Model
	ProductID uint     // 商品ID
	Product   *Product // 商品信息
	Specs     string   // 规格信息，JSON对象
	Price     int64    // SKU价格
	Stock     int32    // 库存数量
	Code      string   `gorm:"uniqueIndex"` // SKU编码
	Version   int32    // 版本号，用于乐观锁
}

// Category 商品分类表
type Category struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         // 分类名称
	ParentID  int64          // 父分类ID，0表示一级分类
	Level     int32          // 分类层级：1-一级分类，2-二级分类，3-三级分类
	SortOrder int32          // 排序序号
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 删除时间
}
