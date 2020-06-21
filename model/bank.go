package model

type Bank struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" gorm:"not null"`
	SWIFTCode SWIFTCode `json:"swift_code" gorm:"not null" validate:"min=8, max=11"`
	Address   string    `json:"address"`
}

type Branch struct {
	ID       int64  `json:"bid"`
	BankCode string `json:"bank_code" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
	Location string `json:"location"`
}

// SWIFT code consists of 8-11 characters.
// Bank code(4) + Country Code (2) + Location code (2) + Branch Code (3)
type SWIFTCode struct {
	BankCode     string `json:"bank_code" validate:"required,eq=4"`     // alphabets
	CountryCode  string `json:"country_code" validate:"required,eq=2"`  // alphabets
	LocationCode string `json:"location_code" validate:"required,eq=2"` // alphanumerics
	BranchCode   string `json:"branch_code" validate:"eq=3"`            // alphanumerics
}

// Request

// Response
