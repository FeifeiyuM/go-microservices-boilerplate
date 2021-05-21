package middlewares

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
)

// EchoRecover panic 恢复
func EchoRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if rec := recover(); rec != nil {
				err := errors.New(fmt.Sprintf("%v", rec))
				fmt.Println(fmt.Sprintf("echo recovery error: %s", err.Error()))
			}
		}()
		return next(c)
	}
}
