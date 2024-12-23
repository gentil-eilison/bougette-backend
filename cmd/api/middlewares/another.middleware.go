package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func AnotherMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("we're in another custom middleware")
		return next(context)
	}
}
