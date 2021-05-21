package rest

import (
	"fmt"
	"gmb/pkg/gmberror"
	"gmb/pkg/log"
	"net/http"
	"github.com/labstack/echo/v4"
	"gmb/internal/service"
)

type httpHandler struct {
	logger log.Factory
}

func NewHttpHandler(logger log.Factory) *httpHandler {
	return &httpHandler{
		logger: logger,
	}
}

// say hello
func (h *httpHandler) sayHello(c echo.Context) error {
	name := c.QueryParam("name")
	msg := fmt.Sprintf("Hello %s", name)
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": msg})
}

func (h *httpHandler) register(c echo.Context) error {
	ctx := c.Request().Context()
	type Account struct {
		Name string `json:"name"`
		Gender int8 `json:"gender"`
		Address string `json:"address"`
	}
	req := &Account{}
	if err := c.Bind(req); err != nil {
		return gmberror.InvalidRequest(err)
	}

	err := service.GetAccountSrv().CreateAccount(ctx, req.Name, req.Address, req.Gender)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": "OK"})
}




func (h *httpHandler) Router(r *echo.Group) {
	// say hello
	r.GET("/say-hello", h.sayHello)
	// 注册
	r.POST("/register", h.register)
}