package mymiddleware

import (
	tokenvalidation "calculator/TokenValidation"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type MiddlewareHandler interface {
	SomeMiddleware(next, stop echo.HandlerFunc) echo.HandlerFunc
	SomeErrorHandler(c echo.Context) error
}

type middleware struct {
	code    int
	message string
}

func NewMiddleware() MiddlewareHandler {
	return &middleware{}
}

func (m *middleware) SetError(code int, message string) {
	m.code = code
	m.message = message
}

func (m *middleware) SomeMiddleware(next, stop echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := c.Request()
		headers := req.Header

		apitoken := headers.Get("Authorization")

		err, verify := tokenvalidation.Isvalid(apitoken)
		if err != nil {
			fmt.Println("im in err")
			m.SetError(http.StatusBadRequest, err.Error())
			return stop(c)
		}
		if !verify {
			fmt.Println("im in not verify section")
			m.SetError(401, "Un-Authorized")
			return stop(c)
		}

		return next(c)
	}
}

func (m *middleware) SomeErrorHandler(c echo.Context) error {
	return c.JSON(
		m.code,
		map[string]any{"message": m.message},
	)
}
