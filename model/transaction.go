package model

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID            uuid.UUID `json:"tid"`
	AccountNumber string    `json:"account_number"` // fk
	CustomerID    string    `json:"customer_id"`    // fk
	Amount        float64   `json:"amount" gorm:"not null"`
	TransferDate  int64     `json:"transfer_date" gorm:"not null"`
	From          string    `json:"from" gorm:"not null"`
	To            string    `json:"to" gorm:"not null"`
	Description   string    `json:"description,omitempty" gorm:"size:150"`
}
