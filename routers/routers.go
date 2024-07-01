package routers

import (
	"calculator/db"
	"calculator/mymiddleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routes(e *echo.Echo) {
	m := mymiddleware.NewMiddleware()
	e.GET("/ws", m.SomeMiddleware(Handlecnnections, m.SomeErrorHandler))
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/migrate", db.Migration)
	e.POST("/user", Crediantials)
	e.POST("/user/login", Login)
	e.POST("/calculator/add", m.SomeMiddleware(Add, m.SomeErrorHandler))
	e.POST("/calculator/substract", m.SomeMiddleware(Substract, m.SomeErrorHandler))
	e.POST("/calculator/multiply", m.SomeMiddleware(Multiply, m.SomeErrorHandler))
	e.POST("/calculator/divide", m.SomeMiddleware(Divide, m.SomeErrorHandler))
	e.GET("/calculator", m.SomeMiddleware(Getall, m.SomeErrorHandler))
	e.GET("/calculator/:id", m.SomeMiddleware(Getbyid, m.SomeErrorHandler))
	e.DELETE("/calculator/:id", m.SomeMiddleware(Delete, m.SomeErrorHandler))
	e.PUT("/calculator/:id", m.SomeMiddleware(Update, m.SomeErrorHandler))
	e.GET("/calculator/symbol", m.SomeMiddleware(Getbysymbol, m.SomeErrorHandler))
	e.POST("/textfileprocessor", m.SomeMiddleware(TextfilePro, m.SomeErrorHandler))
	e.GET("/textfilestats/all", m.SomeMiddleware(Getallstats, m.SomeErrorHandler))

}
