package validator

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
)

var validate *validator.Validate

var validateType = reflect.TypeOf(validator.ValidationErrors{})

const (
	Required = "required"
	Alpha    = "alpha"
	AlphaNum = "alphanum"
	Length   = "len"
)

const (
	StartString = "for '"
	EndString   = "' failed"
)

// NewBank validates input new bank data against a format pattern.
func NewBank(newBank *model.Bank) error {
	validate = validator.New()
	err := validate.Struct(newBank)
	if reflect.TypeOf(err) == validateType {
		eMessage := err.Error()
		fieldName := getFieldName(eMessage)

		r := strings.Contains(eMessage, Required)
		if r {
			return errs.ErrEmptyField(err, fieldName)
		}
		l := strings.Contains(eMessage, Length)
		if l {
			return errs.ErrFieldLength(err, fieldName)
		}
		a := strings.Contains(eMessage, Alpha)
		an := strings.Contains(eMessage, AlphaNum)
		if a || an {
			return errs.ErrDataType(err, fieldName)
		}
	}
	// Todo : check Address
	return nil
}

// getFieldName gets struct's specific field name.
func getFieldName(errorMessage string) (fieldName string) {
	start := strings.Index(errorMessage, StartString)
	end := strings.Index(errorMessage, EndString)
	fieldName = errorMessage[start+len(StartString) : end]
	return fieldName
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
