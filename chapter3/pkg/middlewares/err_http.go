package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
)


type httpErrRet struct {
	Code    string      `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

// EchoErrorHandler 错误拦截中间件
func EchoErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// TODO 待完善错误处理
			fmt.Println(fmt.Sprintf("http error: %s", err.Error()))
		}
		return err
	}
}
