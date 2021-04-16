package middlewares

import (
	"bufio"
	"bytes"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)


type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func httpStdLog(c echo.Context, reqBody, respBody []byte, start time.Time) {
	// TODO 待完善日志输出
	return
}

// EchoStandardLogger echo 标准日志输出
func EchoStandardLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// 请求参数拷贝
		reqBody := []byte{}
		if c.Request().Body != nil {
			reqBody, _ = ioutil.ReadAll(c.Request().Body)
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		// 返回结果拷贝
		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		err := next(c)
		// 输出日志
		httpStdLog(c, reqBody, resBody.Bytes(), start)
		return err
	}
}
