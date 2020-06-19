package server

import (
	"fmt"
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

var d *gorm.DB

func Run() {
	e := NewEcho()
	apiPath := os.Getenv("WELLS_FAR_GO_CONFIG")
	if err := config.InitAPIConfig(apiPath); err != nil {
		log.Fatalf("init api config : '%v'", err)
	}

	db := NewDB()
	// Todo : Create Table
	if err := createTable(db); err != nil {
		log.Fatalf("create database table : '%v'", err)
	}

	e.Logger.Fatal(e.Start(api.PORT))
}

func (s *Server) createTable() error {
	db := NewDB
	if err := s.Database.AutoMigrate(
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

func NewDB() *gorm.DB {
	con := config.GetConfig()
	dbInfo := con.ConnectDB()
	db, err := gorm.Open(con.DB.Driver, dbInfo)
	if err != nil {
		fmt.Sprintf("err : %s", err)
	}
	db.LogMode(true)
	return db
}

func Conn() *gorm.DB {
	return d
}
