package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"gmb/pkg/mq"
)

type mqHandler struct {}

func NewMqHandler() *mqHandler {
	return &mqHandler{}
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

// register
func (m *mqHandler) Register(c *mq.NsqConsumer) {
	if err := c.RegisterHandler("test-recv_hello", "test", m.recvHello); err != nil {
		panic(err)
	}
}