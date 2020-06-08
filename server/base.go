package server

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/8luebottle/Wells-Far-Go/api"
	mw "github.com/8luebottle/Wells-Far-Go/middleware"
)

type Server struct {
	Database *gorm.DB
}

func Run() {
	e := NewEcho()
	fmt.Print(api.BANNER)
	e.Logger.Fatal(e.Start(api.PORT))
}

func NewEcho() *echo.Echo {
	e := echo.New()
	InitRouters(e)
	mw.EnableCORS(e)
	return e
}
