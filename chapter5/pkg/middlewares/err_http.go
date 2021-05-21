package middlewares

import (
	"github.com/labstack/echo/v4"
	"gmb/pkg/gmberror"
	"net/http"
)

type httpErrRet struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Success   bool   `json:"success"`
	ErrDetail string `json:"err_detail"`
}

func sendWarnMsg(ge error) {
	// TODO send msg
}

// UnCatchErrorHandler 未知错误处理
func UnCatchErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errRet := &httpErrRet{
		Code:      "5000",
		Msg:       "小的尽力了！",
		ErrDetail: err.Error(),
		Success:   false,
	}
	_ = c.JSON(code, errRet)
}

func httpGMBErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	errRet := &httpErrRet{}
	if ge, ok := err.(gmberror.GMBError); ok {
		code = ge.HttpStatus()
		errRet.Code = ge.Code()
		errRet.Msg = ge.Message()
		errRet.ErrDetail = ge.Error()
		if ge.SendMsg() {
			sendWarnMsg(ge)
		}
		_ = c.JSON(code, errRet)
	} else {
		UnCatchErrorHandler(err, c)
	}
}

// EchoErrorHandler 错误拦截中间件
func EchoErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			httpGMBErrorHandler(err, c)
		}
		return err
	}
}
