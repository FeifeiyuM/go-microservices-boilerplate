package middlewares

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// EchoRequestID request id 设置
func EchoRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		rID := r.Header.Get("X-Request-Id")
		if rID == "" {
			uid, _ := uuid.NewUUID()
			rID = uid.String()
		}
		r = r.WithContext(context.WithValue(r.Context(), "_requestID", rID))
		c.SetRequest(r)
		return next(c)
	}
}
