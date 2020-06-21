package errs

import (
	goerrors "errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidContext    = goerrors.New("invalid context")
	ErrWellsFarGoContext = goerrors.New("wells-far-go context error")
	ErrBindRequest       = goerrors.New("can not bind the request")
)

type AppErr struct {
	Err     error
	Message string
	Code    int
}

func NewAppErr(err error, message string, code int) *AppErr {
	return &AppErr{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func ErrEmptyField(err error, fieldName string) *AppErr {
	return &AppErr{
		Err:     err,
		Message: fmt.Sprintf("%s can not be empty", fieldName),
		Code:    http.StatusBadRequest,
	}
}

func ErrFieldLength(err error, fieldName string) *AppErr {
	return &AppErr{
		Err:     err,
		Message: fmt.Sprintf("%s is too short or too long", fieldName),
		Code:    http.StatusBadRequest,
	}
}

func ErrDataType(err error, fieldName string) *AppErr {
	return &AppErr{
		Err:     err,
		Message: fmt.Sprintf("invalid data input for %s", fieldName),
		Code:    http.StatusBadRequest,
	}
}

func (e *AppErr) Unwrap() error {
	return e.Err
}

func (e *AppErr) Error() string {
	return fmt.Sprintf("code: %d | err: %v | message: %s", e.Code, e.Err, e.Message)
}
