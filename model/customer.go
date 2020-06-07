package model

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type Gender int8

// Gender
const (
	Male = iota + 1
	Female
	NotSpecified
)

// Todo : Segment customer status & Country Code
type Customer struct {
	ID             string  `json:"cid" gorm:"PRIMARY_KEY"`
	Password       string  `json:"-"`
	Status         int8    `json:"status"`
	FirstName      string  `json:"first_name" gorm:"not null"`
	MiddleName     string  `json:"middle_name"`
	LastName       string  `json:"last_name" gorm:"not null"`
	Gender         Gender  `json:"gender" gorm:"not null"`
	Salary         float64 `json:"salary"`
	CountryCode    int32   `json:"country_code"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email" gorm:"not null"`
	DOB            int32   `json:"date_of_birth"`
	RegisteredDate int64   `json:"registered_date" gorm:"not null"`
	UpdatedAt      int64   `json:"updated_at" gorm:"not null"`
	DeletedAt      int64   `json:"deleted_at,omitempty"`
}

var genderToString = map[Gender]string{
	Male:         "Male",
	Female:       "Female",
	NotSpecified: "Not Specified",
}

var genderToInt = map[string]Gender{
	"Male":          Male,
	"Female":        Female,
	"Not Specified": NotSpecified,
}

func (g *Gender) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(genderToString[g])
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return errors.Wrap(err, "unmarshal gender json")
	}
	*g = genderToInt[j]
	return nil
}
