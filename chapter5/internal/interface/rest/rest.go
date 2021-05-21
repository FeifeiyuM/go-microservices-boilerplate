package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gmb/internal/proto/pb"
	"gmb/internal/service"
	"gmb/pkg/gmberror"
	"gmb/pkg/log"
	"net/http"
)

type httpHandler struct {
	logger log.Factory
}

func NewHttpHandler(logger log.Factory) *httpHandler {
	return &httpHandler{
		logger: logger,
	}
}

func (h *httpHandler) Router(r *echo.Group) {
	// say hello
	r.GET("/say-hello", h.sayHello)
	// 注册
	r.POST("/register", h.register)
	// 充值
	r.POST("/recharge", h.recharge)
}

// say hello
func (h *httpHandler) sayHello(c echo.Context) error {
	name := c.QueryParam("name")
	msg := fmt.Sprintf("Hello %s", name)
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": msg})
}

func (h *httpHandler) register(c echo.Context) error {
	ctx := c.Request().Context()

	req := &pb.RegisterReq{}
	if err := c.Bind(req); err != nil {
		return gmberror.InvalidRequest(err)
	}

	errG := service.GetAccountSrv().CreateAccount(ctx, req.Name, req.Avatar, req.Mobile, int8(req.Gender))
	if errG != nil {
		return errG
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": "OK"})
}

// 充值实现
func (h *httpHandler) recharge(c echo.Context) error {
	ctx := c.Request().Context()
	req := &pb.AccountRechargeReq{}
	if err := c.Bind(req); err != nil {
		return gmberror.InvalidRequest(err)
	}
	errG := service.GetOrderSrv().AccountRecharge(ctx, req.AccId, int64(req.Amount), req.PayOrderId)
	if errG != nil {
		return errG
	}
	return c.JSON(http.StatusOK, &pb.AccountRechargeReply{Message: "OK"})
}
