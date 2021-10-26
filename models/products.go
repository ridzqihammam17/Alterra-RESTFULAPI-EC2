package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID    int    `gorm:"primaryKey" json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
	Stock int    `json:"stock" form:"stock"`

	Carts []*Carts `gorm:"many2many:cart_details" json:"carts"`
	//1 to many with category
	CategoryID int `gorm:"column:category_id" json:"category_id" form:"category_id"`
	Category   Category
}

type GormProductModel struct {
	db *gorm.DB
}

func NewProductModel(db *gorm.DB) *GormProductModel {
	return &GormProductModel{db: db}
}

// Interface Product
type ProductModel interface {
	GetAll() ([]Product, error)
	Get(productId int) (Product, error)
	Insert(Product) (Product, error)
	Edit(product Product, productId int) (Product, error)
	Delete(productId int) (Product, error)
	CheckProductId(productId int) (interface{}, error)
}

func (m *GormProductModel) CheckProductId(productId int) (interface{}, error) {
	var product []Product
	if err := m.db.Where("id=?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (m *GormProductModel) GetAll() ([]Product, error) {
	var product []Product
	if err := m.db.Preload("Category").Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (m *GormProductModel) Get(productId int) (Product, error) {
	var product Product
	if err := m.db.Preload("Category").Where("id=?", productId).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Insert(product Product) (Product, error) {
	if err := m.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Edit(newProduct Product, productId int) (Product, error) {
	var product Product
	var category Category

	if err := m.db.Where("id=?", newProduct.CategoryID).First(&category).Error; err != nil {
		return product, err
	}

	if err := m.db.Find(&product, productId).Error; err != nil {
		return product, err
	}

	product.Name = newProduct.Name
	product.Price = newProduct.Price
	product.Stock = newProduct.Stock
	product.CategoryID = newProduct.CategoryID

	if err := m.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (m *GormProductModel) Delete(productId int) (Product, error) {
	var product Product
	if err := m.db.Find(&product, productId).Error; err != nil {
		return product, err
	}

	if err := m.db.Delete(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
