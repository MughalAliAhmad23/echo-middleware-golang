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
		//AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	routers.Routes(e)
	e.Logger.Fatal(e.Start(":1323"))
	//database k koi b kam krna hai to db wala structure use ho ga chahay data dalna ho ya data nikalna ho or phir agr user to user ko show krna ha to same sturcture use ho ga without db wala or wo structure response k ek general sturcture ma jay ga or wo user ko jay ga...
}
