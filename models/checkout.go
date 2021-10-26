package models

import (
	"time"

	"gorm.io/gorm"
)

type Checkout struct {
	ID        int `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
}

type Checkout_Response struct {
	ID      int
	Product []Checkout_Response
}

type Checkout_Input struct {
	Courier   string `json:"courier" form:"courier"`
	ProductID []int  `json:"product_id" form:"product_id"`
}

type GormCheckoutModel struct {
	db *gorm.DB
}

func NewCheckoutModel(db *gorm.DB) *GormCheckoutModel {
	return &GormCheckoutModel{db: db}
}

type CheckoutModel interface {
	AddCheckoutID() (Checkout, error)
}

func (m *GormProductModel) AddCheckoutID() (Checkout, error) {
	var checkout Checkout
	if err := m.db.Save(&checkout).Error; err != nil {
		checkout := new(Checkout)
		checkout.ID = 0
		return *checkout, nil
	}

	return checkout, nil
}
