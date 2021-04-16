package middlewares

import (
	"context"

	"github.com/nsqio/go-nsq"
	"github.com/opentracing/opentracing-go"

	"gmb/pkg/mq"
)

// MqTracer 链路追踪
func MqTracer(handler mq.MqHandlerFunc) mq.MqHandlerFunc {
	return func(ctx context.Context, msg *nsq.Message) error {
		tr := opentracing.GlobalTracer()
		span := tr.StartSpan("MQ")
		span.SetTag("handler", handler)
		body := msg.Body
		if len(body) > 1024 {
			body = body[:1024]
		}
		span.SetBaggageItem("message", string(body))
		ctx = opentracing.ContextWithSpan(ctx, span)
		defer func() {
			span.Finish()
		}()
		return handler(ctx, msg)
	}
}
