package model

type Bank struct {
	BankCode string `json:"bank_code" gorm:"PRIMARY_KEY"`
	Name     string `json:"name" gorm:"not null"`
	Address  string `json:"address"`
}

type Branch struct {
	ID       int64  `json:"bid"`
	BankCode string `json:"bank_code" gorm:"not null"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
