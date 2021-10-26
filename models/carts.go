package models

import (
	"gorm.io/gorm"
)

type Carts struct {
	gorm.Model
	ID                 int    `gorm:"primaryKey" json:"id" form:"id"`
	StatusTransactions string `json:"status_transactions" form:"status_transactions"`
	TotalQuantity      int    `json:"total_quantity" form:"total_quantity"`
	TotalPrice         int    `json:"total_price" form:"total_price"`

	//many to many
	Products []*Product `gorm:"many2many:cart_details" json:"products"`

	//1 to many
	CustomersID      int `json:"customers_id" form:"customers_id"`
	PaymentMethodsID int `json:"payment_methods_id" form:"payment_methods_id"`

	TransactionID int `json:"transactions_id" form:"transactions_id"`
}

type GormCartsModel struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) *GormCartsModel {
	return &GormCartsModel{db: db}
}

type CartModel interface {
	CreateCart(cart Carts) (Carts, error)
	GetCart(cartId int) (Carts, error)
	GetTotalPrice(cartId int) (int, error)
	GetTotalQty(cartId int) (int, error)
	UpdateTotalCart(cartId int, newTotalPrice int, newTotalQty int) (Carts, error)
	CheckCartId(cartId int) (interface{}, error)
	GetCartById(id int) (Carts, error)
	DeleteCart(cartId int) (cart Carts, err error)
}

func (m *GormCartsModel) CreateCart(cart Carts) (Carts, error) {
	if err := m.db.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

// get cart by id
func (m *GormCartsModel) GetCart(cartId int) (Carts, error) {
	var cart Carts
	if err := m.db.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

// get total price
func (m *GormCartsModel) GetTotalPrice(cartId int) (int, error) {
	var totalPrice int
	if err := m.db.Select("sum(cart_details.price*cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalPrice).Error; err == nil {
		return totalPrice, err
	}
	return totalPrice, nil
}

//get total quantity
func (m *GormCartsModel) GetTotalQty(cartId int) (int, error) {
	var cartDetails CartDetails
	var totalQty int
	if err := m.db.Model(&cartDetails).Select("SUM(cart_details.quantity)").Joins("JOIN carts ON carts.id = cart_details.carts_id").Where("carts_id=?", cartId).First(&totalQty).Error; err == nil {
		return totalQty, err
	}
	return totalQty, nil
}

//update total cart
func (m *GormCartsModel) UpdateTotalCart(cartId int, newTotalPrice int, newTotalQty int) (Carts, error) {
	var cart Carts

	if err := m.db.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}

	cart.TotalQuantity += newTotalQty
	cart.TotalPrice += (newTotalPrice * newTotalQty)

	if err := m.db.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

//check is cart id exist on table cart
func (m *GormCartsModel) CheckCartId(cartId int) (interface{}, error) {
	var cart []Carts
	if err := m.db.Where("id=?", cartId).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// get cart by id
func (m *GormCartsModel) GetCartById(id int) (Carts, error) {
	var cart Carts
	if err := m.db.Find(&cart, "id=?", id).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

//delete cart
func (m *GormCartsModel) DeleteCart(cartId int) (cart Carts, err error) {

	if err := m.db.Find(&cart, "id=?", cartId).Unscoped().Delete(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
