package routers

import (
	"calculator/mymiddleware"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {

	m := mymiddleware.NewMiddleware()
	e.POST("/user", Crediantials)
	e.POST("/user/login", Login)
	e.POST("/calculator/add", m.SomeMiddleware(Add, m.SomeErrorHandler))
	e.POST("/calculator/substract", m.SomeMiddleware(Substract, m.SomeErrorHandler))
	e.POST("/calculator/multiply", m.SomeMiddleware(Multiply, m.SomeErrorHandler))
	e.POST("/calculator/divide", m.SomeMiddleware(Division, m.SomeErrorHandler))
	e.GET("/calculator", m.SomeMiddleware(Getall, m.SomeErrorHandler))
	e.GET("/calculator/:id", m.SomeMiddleware(Getbyid, m.SomeErrorHandler))
	e.DELETE("/calculator/:id", m.SomeMiddleware(Delete, m.SomeErrorHandler))
	e.PUT("/calculator/:id", m.SomeMiddleware(Update, m.SomeErrorHandler))
	e.GET("/calculator/symbol/:operation", m.SomeMiddleware(Getbysymbol, m.SomeErrorHandler))
}
