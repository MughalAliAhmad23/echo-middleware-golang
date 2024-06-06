package routers

import (
	"calculator/mymiddleware"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {

	m := mymiddleware.NewMiddleware()
	e.POST("/user", Crediantials)
	e.POST("/user/login", Login) //handler change nai ho ga or route ma hi middle ware k through token check ho ga or add k func call honay se phlay phlay tk
	//ya check krna hai k post man se token kasay set krty hain
	e.POST("/calculator/add", m.SomeMiddleware(Add, m.SomeErrorHandler))
	e.POST("/calculator/substract", Substract)
	e.POST("/calculator/multiply", Multiply)
	e.POST("/calculator/divide", Division)
	e.GET("/calculator", Getall)
	e.GET("/calculator/:id", Getbyid)
	e.DELETE("/calculator/:id", Delete)
	e.PUT("/calculator/:id", Update)
	e.GET("/calculator/symbol/:operation", Getbysymbol)
}
