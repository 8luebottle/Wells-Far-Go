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
	Password       string  `json:"password" validate:"min=6,max=12"` // Todo : Add Regex
	Status         int8    `json:"status"`
	FirstName      string  `json:"first_name" gorm:"not null" validate:"alphaunicode"`
	MiddleName     string  `json:"middle_name" gorm:"omitempty" validate:"alphaunicode"`
	LastName       string  `json:"last_name" gorm:"not null" validate:"alphaunicode"`
	Gender         Gender  `json:"gender" gorm:"not null"`
	Salary         float64 `json:"salary" gorm:"omitempty"`
	CountryCode    int32   `json:"country_code" gorm:"not null" validate:"min=1,max=5"`
	Phone          string  `json:"phone" validate:"required,numeric"` // no '-'
	Email          string  `json:"email" gorm:"not null" validate:"required"`
	DOB            int32   `json:"date_of_birth" validate:"required,len=8"`
	RegisteredDate int64   `json:"registered_date" gorm:"not null"`
	UpdatedAt      int64   `json:"updated_at" gorm:"not null"`
	DeletedAt      int64   `json:"deleted_at,omitempty"`
}

func (c *Customer) MarshalJSON() ([]byte, error) {
	r := &customerResponse{
		ID:             c.ID,
		Status:         c.Status,
		FirstName:      c.FirstName,
		MiddleName:     c.MiddleName,
		LastName:       c.LastName,
		Gender:         c.Gender,
		Phone:          c.Phone,
		Email:          c.Email,
		DOB:            c.DOB,
		RegisteredDate: c.RegisteredDate,
	}
	return json.Marshal(r)
}

type customerResponse struct {
	ID             string `json:"cid"`
	Status         int8   `json:"status"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	Gender         Gender `json:"gender"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	DOB            int32  `json:"dob"`
	RegisteredDate int64  `json:"registered_date"`
}

// map[Gender]string
var genderToString = map[interface{}]string{
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
