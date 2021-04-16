package middlewares

import (
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// GrpcUnaryTracerInterceptor grpc tracer
func GrpcUnaryTracerInterceptor() grpc.UnaryServerInterceptor {
	tr := opentracing.GlobalTracer()
	return grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tr))
}
