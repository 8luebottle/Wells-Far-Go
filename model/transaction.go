package model

import (
	"github.com/google/uuid"
)

type TransactionType string

// Transaction Type
const (
	Authorization = "authorization"
	Credit        = "credit"
	Debit         = "debit"
	Refund        = "refund"
	Void          = "void"
)

type TransactionStatus string

// Transaction Status
const (
	Approved  = "approved"
	Pending   = "pending"
	Declined  = "declined"
	Voided    = "voided"
	Suspended = "suspend"
)

type Transaction struct {
	ID            uuid.UUID         `json:"tid"`
	AccountNumber string            `json:"account_number"` // fk
	CustomerID    string            `json:"cid"`            // fk
	Type          TransactionType   `json:"type"`
	Status        TransactionStatus `json:"transaction_status"`
	Amount        float64           `json:"amount" gorm:"not null"`
	TransferDate  int64             `json:"transfer_date" gorm:"not null"`
	From          string            `json:"from" gorm:"not null"`
	To            string            `json:"to" gorm:"not null"`
	Description   string            `json:"description,omitempty" gorm:"size:150"`
}
