package middlewares

import (
	"context"
	"fmt"

	"github.com/nsqio/go-nsq"

	"gmb/pkg/mq"
)

// MqErrHandler nsq 错误处理
func MqErrHandler(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) error {
		err := handler(ctx, msg)
		if err != nil {
			// TODO 待完善错误处理
			fmt.Println(fmt.Sprintf("mq error: %s", err.Error()))
		}
		return err
	}
}
