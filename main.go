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
	//database k koi b kam krna hai to db wala structure use ho ga chahay data dalna ho ya data nikalna ho or phir agr user to user ko show krna ha to same sturcture use ho ga without db wala or wo structure response k ek general sturcture ma jay ga or wo user ko jay ga...
}
