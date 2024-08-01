package main

import (
	"calculator/db"
	"calculator/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Swagger Example Api
// @version 1.0
// @description This is sample calculator & user server
// @host localhost:1323
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**
func main() {

	e := echo.New()
	db.Connect()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	routers.Routes(e)
	e.Logger.Fatal(e.Start(":1323"))

}
