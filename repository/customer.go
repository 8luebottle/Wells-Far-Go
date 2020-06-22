package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/8luebottle/Wells-Far-Go/model"
)

type CustomerStorer interface {
	Create(db *gorm.DB, customer *model.Customer) error
}

type customerRepository struct{}

func ParseCustomerStorer() CustomerStorer {
	return &customerRepository{}
}

func (cr *customerRepository) Create(db *gorm.DB, customer *model.Customer) error {
	if err := db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}
