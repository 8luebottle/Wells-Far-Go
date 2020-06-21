package validator

import (
	"errors"
	"strings"

	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
)

// Validates Bank
func NewBank(bank *model.Bank) error {
	// check Name not null
	if err := checkEmptyString(
		bank.Name, bank.Address,
		bank.SWIFTCode.BankCode,
		bank.SWIFTCode.CountryCode,
		bank.SWIFTCode.LocationCode,
	); err != nil {
		return errs.EmptyField(err)
	}
	// check SWIFT CODE
	// check Address
	return nil
}

// checkEmptyString checks empty string fields.
func checkEmptyString(stringFields ...string) error {
	for _, field := range stringFields {
		f := strings.TrimSpace(field)
		if len(f) == 0 {
			return errors.New("field can not be empty")
		}
	}
	return nil
}
