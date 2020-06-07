package model

type AddressType int8

// Address Type
const (
	Home = iota + 1
	Work
	Mailing
)

type Address struct {
	ID            int64       `json:"aid"`
	CustomerID    string      `json:"cid"`            // fk
	AccountNumber string      `json:"account_number"` // fk
	Type          AddressType `json:"type"`
	State         string      `json:"state"`
	City          string      `json:"city"`
	Street        string      `json:"street"`
	ZipCode       int64       `json:"zip_code"`
	Country       string      `json:"country"`
}
