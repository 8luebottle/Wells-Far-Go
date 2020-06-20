package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
)

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
