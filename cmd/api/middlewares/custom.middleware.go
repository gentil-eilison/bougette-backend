package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("we're in the custom middleware")
		context.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(context)
	}
}
