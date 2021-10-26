package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Order  string `json:"order_id" form:"order_id"`
	Amount int    `json:"gross_amount" form:"gross_amount"`
}

type GormPaymentModel struct {
	db *gorm.DB
}

func NewPaymentModel(db *gorm.DB) *GormPaymentModel {
	return &GormPaymentModel{db: db}
}

// Interface Payment
type PaymentModel interface {
	Get(paymentId int) (Payment, error)
	Add(Payment) (Payment, error)
}

func (m *GormPaymentModel) Get(paymentId int) (Payment, error) {
	var payment Payment
	if err := m.db.Where("id=?", paymentId).First(&payment).Error; err != nil {
		return payment, err
	}
	return payment, nil
}

func (m *GormPaymentModel) Add(payment Payment) (Payment, error) {
	return payment, nil
}
