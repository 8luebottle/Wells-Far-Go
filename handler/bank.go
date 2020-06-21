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

type BankController interface {
	CreateBank(c echo.Context) error
}

type bank struct {
	bs service.BankServer
}

func ParseBankController(bs service.BankServer) BankController {
	return &bank{bs: bs}
}

func (b *bank) CreateBank(c echo.Context) error {
	ctx, err := middleware.WellsFarGoContext(c)
	if err != nil {
		return errors.Wrapf(errs.ErrWellsFarGoContext, "err : '%s'", err)
	}

	var newBank model.Bank
	if err = ctx.Bind(&newBank); err != nil {
		return errors.Wrapf(errs.ErrBindRequest, "err : '%s'", err)
	}

	bank, err := b.bs.CreateNewBank(&newBank)
	if err != nil {
		return errors.Wrap(err, "create a new bank")
	}

	return ctx.JSON(http.StatusOK, bank)
}

/*
	Implement List
		2. Create Branch of the Bank
		3. Update Bank Information
		4. Delete Bank (Soft Delete)
		5. Get List of Bank
*/
