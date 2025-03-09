package model

import (
	"time"

	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	ID          uint32    `gorm:"primary_key;auto_increment"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Picture     string    `gorm:"type:varchar(255)"`
	Price       float32   `gorm:"type:decimal(10,2);not null"`
	Categories  string    `gorm:"type:text"` // 使用JSON字符串存储分类数组
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName 设置表名
func (Product) TableName() string {
	return "products"
}

// CreateProduct 创建商品
func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

// GetProductByID 通过ID获取商品
func GetProductByID(db *gorm.DB, id uint32) (*Product, error) {
	var product Product
	err := db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct 更新商品
func UpdateProduct(db *gorm.DB, product *Product) error {
	return db.Save(product).Error
}

// DeleteProduct 删除商品
func DeleteProduct(db *gorm.DB, id uint32) error {
	return db.Delete(&Product{}, id).Error
}

// ListProducts 列出商品
func ListProducts(db *gorm.DB, offset, limit int, category string) ([]Product, int64, error) {
	var products []Product
	var total int64

	query := db
	if category != "" {
		query = query.Where("JSON_CONTAINS(categories, JSON_ARRAY(?))", category)
	}

	if err := query.Model(&Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// SearchProducts 搜索商品
func SearchProducts(db *gorm.DB, keywords string, offset, limit int) ([]Product, int64, error) {
	var products []Product
	var total int64

	query := db.Where("name LIKE ? OR description LIKE ?", "%"+keywords+"%", "%"+keywords+"%")

	if err := query.Model(&Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
