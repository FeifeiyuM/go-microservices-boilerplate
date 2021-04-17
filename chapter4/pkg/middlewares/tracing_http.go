package middlewares

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
)

// TracingInit 初始化 jeager
func TracingInit(serviceName string, metricsFactory metrics.Factory) (opentracing.Tracer, io.Closer) {
	cfg, _ := config.FromEnv()
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	metricsFactory = metricsFactory.Namespace(metrics.NSOptions{Name: serviceName, Tags: nil})
	tracer, closer, _ := cfg.NewTracer(
		//config.Logger(jaegerLogger),
		config.Metrics(metricsFactory),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)

	// set the singleton operation.Tracer with the jaeger tracer
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

// EchoTracer echo jeager 中间件
func EchoTracer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		// tracing
		tr := opentracing.GlobalTracer()
		ctx, _ := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := tr.StartSpan("HTTP "+r.Method+" "+r.URL.Path, ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.String())
		ext.Component.Set(sp, "net/http")

		defer func() {
			ext.HTTPStatusCode.Set(sp, uint16(c.Response().Status))
			if c.Response().Status > http.StatusNotModified {
				ext.Error.Set(sp, true)
			}
			sp.Finish()
		}()
		return next(c)
	}
}
