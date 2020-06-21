package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/8luebottle/Wells-Far-Go/model"
)

type BankStorer interface {
	Create(db *gorm.DB, bank *model.Bank) error
}

type bankRepository struct{}

func ParseBankStorer() BankStorer {
	return &bankRepository{}
}

func (br *bankRepository) Create(db *gorm.DB, bank *model.Bank) error {
	if err := db.Create(bank).Error; err != nil {
		return err
	}
	return nil
}
