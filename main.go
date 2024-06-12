package main

import (
	"calculator/db"
	"calculator/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db.Connect()
	routers.Routes(e)
	e.Logger.Fatal(e.Start(":1323"))

}
