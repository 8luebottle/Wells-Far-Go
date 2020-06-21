package model

type Bank struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	SWIFTCode SWIFTCode `json:"swift_code" gorm:"not null" validate:"unique"`
	Address   string    `json:"address" validate:"required"`
}

type Branch struct {
	ID       int64  `json:"bid"`
	BankCode string `json:"bank_code" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Location string `json:"location"`
}

// SWIFT code consists of 8-11 characters.
// Bank code(4) + Country Code (2) + Location code (2) + Branch Code (3) optional
type SWIFTCode struct {
	BankCode     string `json:"bank_code" validate:"required,len=4,alpha"`        // alphabets
	CountryCode  string `json:"country_code" validate:"required,len=2,alpha"`     // alphabets
	LocationCode string `json:"location_code" validate:"required,len=2,alphanum"` // alphanumerics
	BranchCode   string `json:"branch_code" validate:"omitempty,alpha"`           // alphanumerics
}

// Request

// Response
