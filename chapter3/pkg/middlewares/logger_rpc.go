package middlewares

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func rpcStdLogger(req, resp []byte, st *status.Status,  start time.Time) {
	// TODO 待完善日志输出
	return
}

// GrpcUnaryStdLoggerInterceptor grpc standard log
func GrpcUnaryStdLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		// 日志整理
		s, _ := status.FromError(err)
		reqB, _ := json.Marshal(req)
		respB, _ := json.Marshal(resp)
		rpcStdLogger(reqB, respB, s, start)
		return
	}
}
