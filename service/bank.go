package service

import (
	"github.com/pkg/errors"

	c "github.com/8luebottle/Wells-Far-Go/config"
	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/repository"
	v "github.com/8luebottle/Wells-Far-Go/validator"
)

type BankServer interface {
	CreateNewBank(newBank *model.Bank) (*model.Bank, error)
}

type bankService struct {
	br repository.BankStorer
}

func ParseBankServer(br repository.BankStorer) BankServer {
	return &bankService{br: br}
}

// CreateNewBank creates new bank.
func (bs *bankService) CreateNewBank(newBank *model.Bank) (*model.Bank, error) {

	if err := v.NewBank(newBank); err != nil {
		return nil, errors.Wrap(err, "validate new bank")
	}

	if err := bs.br.Create(c.DBConn(), newBank); err != nil {
		return nil, errors.Wrap(err, "create new bank to db table")
	}

	return newBank, nil
}

// GetSWIFTCode retrieves SWIFT code of the specific bank.
func (bs *bankService) GetSWIFTCode(bank *model.Bank) (string, error) {
	// Todo : Search from DB
	SWIFTCode := bank.SWIFTCode.BankCode +
		bank.SWIFTCode.CountryCode +
		bank.SWIFTCode.LocationCode +
		bank.SWIFTCode.BranchCode
	return SWIFTCode, nil
}
