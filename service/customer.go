package service

import (
	"github.com/pkg/errors"

	c "github.com/8luebottle/Wells-Far-Go/config"
	"github.com/8luebottle/Wells-Far-Go/model"
	"github.com/8luebottle/Wells-Far-Go/repository"
	v "github.com/8luebottle/Wells-Far-Go/validator"
)

type CustomerServer interface {
	CreateNewCustomer(newCustomer *model.Customer) (*model.Customer, error)
}

type customerService struct {
	cr repository.CustomerStorer
}

func ParseCustomerServer(cr repository.CustomerStorer) CustomerServer {
	return &customerService{cr: cr}
}

func (cs *customerService) CreateNewCustomer(newCustomer *model.Customer) (*model.Customer, error) {
	if err := v.NewRequest(newCustomer); err != nil {
		return nil, errors.Wrap(err, "validate new customer")
	}

	if err := cs.cr.Create(c.DBConn(), newCustomer); err != nil {
		return nil, errors.Wrap(err, "create new customer to db table")
	}

	return newCustomer, nil
}
