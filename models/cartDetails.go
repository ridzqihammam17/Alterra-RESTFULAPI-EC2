package models

import (
	"time"

	"gorm.io/gorm"
)

type CartDetails struct {
	ProductsID int `gorm:"primaryKey" json:"products_id" form:"products_id"`
	CartsID    int `gorm:"primaryKey" json:"carts_id" form:"carts_id"`
	Quantity   int `json:"quantity" form:"quantity"`
	Price      int `json:"price" form:"price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type GormCartDetailsModel struct {
	db *gorm.DB
}

func NewCartDetailModel(db *gorm.DB) *GormCartDetailsModel {
	return &GormCartDetailsModel{db: db}
}

type CartDetailModel interface {
	CheckProductAndCartId(productId, cartId int, cartDetails CartDetails) (interface{}, error)
	GetCartDetailByCartId(cartId int) (CartDetails, error)
	AddToCart(cartDetails CartDetails) (CartDetails, error)
	DeleteProductFromCart(cartId, productId int) (interface{}, error)
	GetListProductCart(cartId int) (interface{}, error)
	CountProductOnCart(cartId int) (int, error)
	CountProductandPriceOnCart(cartId int) (int, int, error)
}

func (m *GormCartDetailsModel) CheckProductAndCartId(productId, cartId int, cartDetails CartDetails) (interface{}, error) {
	if err := m.db.Where("carts_id=? AND products_id=?", cartId, productId).First(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}

// get product by id
func (m *GormCartDetailsModel) GetProduct(productId int) (Product, error) {
	var product Product
	if err := m.db.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}
	return product, nil
}

//Get cart details by Cart ID
func (m *GormCartDetailsModel) GetCartDetailByCartId(cartId int) (CartDetails, error) {
	var cartDetails CartDetails
	if err := m.db.Find(&cartDetails, "carts_id=?", cartId).Error; err != nil {
		return cartDetails, err
	}
	return cartDetails, nil
}

//add product to cart
func (m *GormCartDetailsModel) AddToCart(cartDetails CartDetails) (CartDetails, error) {
	if err := m.db.Save(&cartDetails).Error; err != nil {
		return cartDetails, err
	}
	return cartDetails, nil
}

//delete product from cart detail
func (m *GormCartDetailsModel) DeleteProductFromCart(cartId, productId int) (interface{}, error) {
	var cartDetails CartDetails
	if err := m.db.Find(&cartDetails, "carts_id=? AND products_id=?", cartId, productId).Unscoped().Delete(&cartDetails).Error; err != nil {
		return nil, err
	}
	return cartDetails, nil
}

//get all products from cart detail
func (m *GormCartDetailsModel) GetListProductCart(cartId int) (interface{}, error) {
	var cartDetail CartDetails

	if err := m.db.Find(&cartDetail, "carts_id=?", cartId).Error; err != nil {
		return nil, err
	}
	// if err := m.db.Table("products").Joins("JOIN cart_details ON products.id = cart_details.products_id").Joins("JOIN carts ON cart_details.carts_id = carts.id").Where("carts.id=?", cartId).Find(&cartDetail).Error; err != nil {
	// 	return cartDetail, nil
	// }
	return cartDetail, nil
}

func (m *GormCartDetailsModel) CountProductOnCart(cartId int) (int, error) {
	var countProduct int
	if err := m.db.Select("COUNT(carts_id)").Where("carts_id=?", cartId).First(&countProduct).Error; err == nil {
		return countProduct, err
	}
	return countProduct, nil
}

func (m *GormCartDetailsModel) CountProductandPriceOnCart(cartId int) (int, int, error) {
	var countProduct, Price int
	if err := m.db.Select("COUNT(carts_id), Price").Where("carts_id=?", cartId).First(&countProduct).Error; err == nil {
		return countProduct, Price, err
	}
	return countProduct, Price, nil
}
