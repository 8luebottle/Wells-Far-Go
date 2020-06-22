package validator

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
)

var (
	validate     *validator.Validate
	validateType = reflect.TypeOf(validator.ValidationErrors{})
)

// Tag of gopkg validator
const (
	Alpha    = "alpha"
	AlphaNum = "alphanum"
	Length   = "len"
	Num      = "numeric"
	Required = "required"
)

const (
	StartString = "for '"
	EndString   = "' failed"
)

type ValidRequest struct {
	fieldName  string
	eMessage   string
	validError error
}

// NewRequest validates new request's fields.
func NewRequest(newStruct interface{}) error {
	validate = validator.New()
	err := validate.Struct(newStruct)
	if reflect.TypeOf(err) == validateType {
		eMessage := err.Error()
		fieldName := getFieldName(eMessage)
		newValidator := &ValidRequest{
			fieldName:  fieldName,
			eMessage:   eMessage,
			validError: err,
		}
		if err := getErrorByTags(newValidator); err != nil {
			return err
		}
	}
	return nil
}

// getFieldName gets struct's specific field name.
func getFieldName(errorMessage string) (fieldName string) {
	start := strings.Index(errorMessage, StartString)
	end := strings.Index(errorMessage, EndString)
	fieldName = errorMessage[start+len(StartString) : end]
	return fieldName
}

// getErrorByTags extracts error message from validator pkg.
func getErrorByTags(vr *ValidRequest) error {
	a := strings.Contains(vr.eMessage, Alpha)
	an := strings.Contains(vr.eMessage, AlphaNum)
	n := strings.Contains(vr.eMessage, Num)
	l := strings.Contains(vr.eMessage, Length)
	r := strings.Contains(vr.eMessage, Required)

	switch {
	case a:
		fallthrough
	case an:
		fallthrough
	case n:
		return errs.ErrDataType(vr.validError, vr.fieldName)
	case l:
		return errs.ErrFieldLength(vr.validError, vr.fieldName)
	case r:
		return errs.ErrEmptyField(vr.validError, vr.fieldName)
	default:
		return nil
	}
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
