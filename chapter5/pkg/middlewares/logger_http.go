package middlewares

import (
	"bufio"
	"bytes"
	"github.com/labstack/echo/v4"
	"gmb/pkg/log"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	logger log.Factory
)

// 初始中间件日志
func InitLogger(l log.Factory) {
	logger = l
}

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
	du := time.Since(start)
	r := c.Request()
	// filter health check log
	if strings.HasPrefix(r.RequestURI, "/ping") {
		return
	}

	fields := make([]zap.Field, 9)
	fields[0] = zap.Int("status_code", c.Response().Status)
	fields[1] = zap.String("duration", du.String())
	fields[2] = zap.String("method", r.Method)
	fields[3] = zap.String("uri", r.RequestURI)
	fields[4] = zap.String("agent", r.UserAgent())
	fields[5] = zap.String("remote_ip", c.RealIP())
	fields[6] = zap.String("token", r.Header.Get("token"))
	fields[7] = zap.ByteString("request_body", reqBody)

	if c.Response().Status > http.StatusNotModified {
		fields[8] = zap.ByteString("response_body", respBody)
		logger.For(r.Context()).Error("failed", fields...)
	} else {
		fields[8] = zap.ByteString("response_body", []byte(""))
		logger.For(r.Context()).Info("success", fields...)
	}
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
