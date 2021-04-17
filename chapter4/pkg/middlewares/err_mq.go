package middlewares

import (
	"context"
	"github.com/nsqio/go-nsq"
	"gmb/pkg/gmberror"

	"gmb/pkg/mq"
)

// MqErrHandler nsq 错误处理
func MqErrHandler(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) error {
		err := handler(ctx, msg)
		if err != nil {
			// 是否发送告警
			if ge, ok := err.(gmberror.GMBError); ok {
				sendWarnMsg(ge)
			} else {
				sendWarnMsg(err)
			}
		}
		return err
	}
}
