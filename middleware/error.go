package middleware

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
)

var validateType = reflect.TypeOf(validator.ValidationErrors{})

type WellsFarGoError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, err := WellsFarGoContext(c)
		if err != nil {
			return errors.Wrapf(errs.ErrWellsFarGoContext, "err : '%s'", err)
		}
		if err = next(ctx); err == nil {
			return nil
		}
		oerr := errors.Cause(err)
		if reflect.TypeOf(err) == validateType {
			return ctx.JSON(200, WellsFarGoError{Message: "validate failed", Code: 400})
		}

		switch err := oerr.(type) {
		case *echo.HTTPError:
			return ctx.JSON(err.Code, WellsFarGoError{Message: err.Message.(string)})
		case *errs.AppErr:
			return ctx.JSON(err.Code, WellsFarGoError{
				Message: err.Message,
				Code:    err.Code,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, WellsFarGoError{
			Message: "wells-far-go internal server error",
			Code:    500,
		})
	}
}
