package model

type AccountType int8

// AccountType
const (
	Checking = iota + 1
	Savings
)

type Account struct {
	Number     int64       `json:"account_number" gorm:"PRIMARY_KEY"`
	CustomerID string      `json:"cid"` // fk
	Type       AccountType `json:"type" gorm:"default:1"`
	Name       string      `json:"name" gorm:"size:50"`
	Balance    float64     `json:"balance" gorm:"not null"`
	OpenedAt   int64       `json:"opened_at" gorm:"not null"`
	DeletedAt  int64       `json:"deleted_at,omitempty"`
}

type Deposit struct {
	AccountNumber string  `json:"account_number"`
	CustomerID    string  `json:"cid"`
	Date          int64   `json:"date" gorm:"not null"`
	Amount        float64 `json:"amount" gorm:"not null"`
}

type Withdraw struct {
	AccountNumber string  `json:"account_number"`
	CustomerID    string  `json:"cid"`
	Date          int64   `json:"date" gorm:"not null"`
	Amount        float64 `json:"amount" gorm:"not null"`
}
