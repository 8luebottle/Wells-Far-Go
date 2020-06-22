package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/8luebottle/Wells-Far-Go/middleware"
	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/pkg/errs"
	"github.com/8luebottle/Wells-Far-Go/service"
)

/*
	Implement List
		2. Create Bank Manager (Employee)
		3. Get New Customer
		4. Update Existence Customer
		5. Delete Customer (Soft delete)
*/

type CustomerController interface {
	CreateCustomer(c echo.Context) error
}

type customer struct {
	cs service.CustomerServer
}

func ParseCustomerController(cs service.CustomerServer) CustomerController {
	return &customer{cs: cs}
}

// Todo : Add Branch Selection
func (ct *customer) CreateCustomer(c echo.Context) error {
	ctx, err := middleware.WellsFarGoContext(c)
	if err != nil {
		return errors.Wrapf(errs.ErrWellsFarGoContext, "err : '%s'", err)
	}

	var newCustomer model.Customer
	if err = ctx.Bind(&newCustomer); err != nil {
		return errors.Wrapf(errs.ErrBindRequest, "err : '%s'", err)
	}

	customer, err := ct.cs.CreateNewCustomer(&newCustomer)
	if err != nil {
		return errors.Wrap(err, "create new customer")
	}

	return ctx.JSON(http.StatusOK, customer)
}

// Admin can create banks, branches of banks.
// Temporarily hold on.
func (ct *customer) CreateAdmin(c echo.Context) error {
	panic("implement me later on")
}
