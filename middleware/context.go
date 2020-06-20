package middleware

import (
	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
	"github.com/labstack/echo/v4"
)

type ContextProvider interface {
	echo.Context

	// Todo : Add Methods
}

type wellsFarGoContext struct {
	echo.Context

	// Todo : Add User Auth
}

func WellsFarGoContext(c echo.Context) (ContextProvider, error) {
	ctx, ok := c.(ContextProvider)
	if !ok {
		return nil, errs.ErrInvalidContext
	}
	return ctx, nil
}
