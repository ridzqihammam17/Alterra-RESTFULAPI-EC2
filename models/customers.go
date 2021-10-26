package models

import (
	"altastore/api/middlewares"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`

	//1 to many with carts
	// Carts []Carts `gorm:"foreignKey:CustomersID"`
	Carts Carts `gorm:"foreignKey:CustomersID"`
}

type GormCustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *GormCustomerModel {
	return &GormCustomerModel{db: db}
}

// Interface Customer
type CustomerModel interface {
	// Get(customerId int) (Customer, error)
	Register(Customer) (Customer, error)
	Login(email, password string) (Customer, error)
	GetAll() ([]Customer, error)
	// Edit(csutomer Customer, customerId int) (Customer, error)
	// Delete(customerId int) (Customer, error)
}

func (m *GormCustomerModel) Register(customer Customer) (Customer, error) {
	// Encrypt Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		return customer, err
	}

	customer.Password = string(hashedPassword)

	if err := m.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Login(email, password string) (Customer, error) {
	var customer Customer
	var err error

	if err = m.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer, err
	}

	// Checking Encrypt Password
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return customer, err
	}

	customer.Token, err = middlewares.CreateToken(int(customer.ID))

	if err != nil {
		return customer, err
	}

	if err := m.db.Save(customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

func (m *GormCustomerModel) GetAll() ([]Customer, error) {
	var customer []Customer
	if err := m.db.Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
