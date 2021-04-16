package middlewares

import (
	"context"
	"time"

	"github.com/nsqio/go-nsq"

	"gmb/pkg/mq"
)

func mqStdLogger(body []byte, du time.Duration) {
	// TODO 待完善日志输出
	return
}

// MqStdLogger nsq 标准日志输出
func MqStdLogger(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) error {
		start := time.Now()
		body := msg.Body
		err := handler(ctx, msg)
		du := time.Since(start)
		// 日志输出
		mqStdLogger(body, du)
		return err
	}
}
