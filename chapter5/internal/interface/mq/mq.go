package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"gmb/internal/proto/pb"
	"gmb/internal/service"
	"gmb/pkg/gmberror"
	"gmb/pkg/log"
	"gmb/pkg/mq"
)

type mqHandler struct {
	logger log.Factory
}

func NewMqHandler(logger log.Factory) *mqHandler {
	return &mqHandler{
		logger: logger,
	}
}

// register
func (m *mqHandler) Register(c *mq.NsqConsumer) {
	if err := c.RegisterHandler("test-recv_hello", "test", m.recvHello); err != nil {
		panic(err)
	}
}

func (m *mqHandler) recvHello(_ context.Context, msg *nsq.Message) error {
	if msg == nil {
		return errors.New("msg is null")
	}
	type param struct {
		Name string `json:"name"`
	}
	req := &param{}
	if err := json.Unmarshal(msg.Body, req); err != nil {
		return err
	}
	fmt.Printf("%s send a hello", req.Name)
	return nil
}

// 充值实现
func (m *mqHandler) recharge(ctx context.Context, msg *nsq.Message) error {
	req := &pb.AccountRechargeReq{}
	err := json.Unmarshal(msg.Body, req)
	if err != nil {
		return gmberror.InvalidRequest(err)
	}
	return service.GetOrderSrv().AccountRecharge(ctx, req.AccId, int64(req.Amount), req.PayOrderId)
}
