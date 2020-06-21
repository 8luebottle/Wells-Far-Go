package service

import (
	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/repository"
)

// Own Bank 는 Wells-Far-Go
type BankServer interface {
	CreateNewBank() (*model.Bank, error)
}

type bankService struct {
	br repository.BankStorer
}

func ParseBankServer(br repository.BankStorer) BankServer {
	return &bankService{br: br}
}

func (bs *bankService) CreateNewBank() (*model.Bank, error) {
	return nil, nil
}
