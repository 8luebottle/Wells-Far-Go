package server

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"github.com/8luebottle/Wells-Far-Go/api"
	"github.com/8luebottle/Wells-Far-Go/config"
	mw "github.com/8luebottle/Wells-Far-Go/middleware"
	"github.com/8luebottle/Wells-Far-Go/model"
)

type Server struct {
	Database *gorm.DB
	Conf     config.Config
}

func Run() {
	e := NewEcho()

	apiPath := os.Getenv(api.APIConfig)
	if err := config.InitAPIConfig(apiPath); err != nil {
		log.Fatalf("init api config : '%v'", err)
	}

	db := config.Conn()
	if err := createTable(db); err != nil {
		log.Fatalf("create database table : '%v'", err)
	}
	e.Logger.Fatal(e.Start(api.PORT))
}

func createTable(d *gorm.DB) error {
	if err := d.AutoMigrate(
		&model.Account{}, &model.Address{},
		&model.Bank{},
		&model.Customer{},
		&model.Transaction{},
	).Error; err != nil {
		return errors.Wrap(err, "auto migrate database table")
	}
	return nil
}

func NewEcho() *echo.Echo {
	e := echo.New()
	InitRouters(e)
	mw.EnableCORS(e)
	return e
}
