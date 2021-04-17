package middlewares

import (
	"context"
	"go.uber.org/zap"
	"time"

	"github.com/nsqio/go-nsq"

	"gmb/pkg/mq"
)

func mqStdLogger(ctx context.Context, body []byte, err error, start time.Time) {
	du := time.Since(start)
	if len(body) > 1024 {
		body = body[:1024]
	}
	if err != nil {
		logger.For(ctx).Error("failed",
			zap.ByteString("message", body),
			zap.String("duration", du.String()),
			zap.String("error", err.Error()))
	} else {
		logger.For(ctx).Info("success",
			zap.ByteString("message", body),
			zap.String("duration", du.String()))
	}
}

// MqStdLogger nsq 标准日志输出
func MqStdLogger(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) error {
		start := time.Now()
		body := msg.Body
		err := handler(ctx, msg)
		// 日志输出
		mqStdLogger(ctx, body, err, start)
		return err
	}
}
