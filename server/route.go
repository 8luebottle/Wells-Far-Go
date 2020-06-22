package server

import (
	"github.com/labstack/echo/v4"

	"github.com/8luebottle/Wells-Far-Go/api"
	"github.com/8luebottle/Wells-Far-Go/handler"
	"github.com/8luebottle/Wells-Far-Go/middleware"
	"github.com/8luebottle/Wells-Far-Go/repository"
	"github.com/8luebottle/Wells-Far-Go/service"
)

func InitRouters(e *echo.Echo) {
	v := e.Group(api.VERSION)
	e.Use(middleware.ErrorMiddleware)

	bankRepo := repository.ParseBankStorer()
	bankSvc := service.ParseBankServer(bankRepo)
	bankHandler := handler.ParseBankController(bankSvc)
	vBank := v.Group("/banks")
	{
		// Todo : Admin Auth
		vBank.POST("/new", bankHandler.CreateBank)
	}

	customerRepo := repository.ParseCustomerStorer()
	customerSvc := service.ParseCustomerServer(customerRepo)
	customerHandler := handler.ParseCustomerController(customerSvc)
	vCustomer := v.Group("/customers")
	{
		// Todo : Customer Auth
		vCustomer.POST("/new", customerHandler.CreateCustomer)
	}
}
