package middlewares

import (
	"context"
	"fmt"

	"github.com/nsqio/go-nsq"

	"gmb/pkg/mq"
)

// MqRecovery panic 恢复
func MqRecovery(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("%v", rec)
				fmt.Println(fmt.Sprintf("mq recovery error: %s", err.Error()))
			}
		}()
		return handler(ctx, msg)
	}
}
