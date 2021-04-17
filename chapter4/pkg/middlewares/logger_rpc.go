package middlewares

import (
	"context"
	"go.uber.org/zap"
	"time"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func rpcStdLogger(ctx context.Context, info *grpc.UnaryServerInfo, req, resp []byte, s *status.Status, err error,  start time.Time) {
	du := time.Since(start)
	fields := make([]zap.Field, 5)
	fields[0] = zap.Int32("rpc_code", int32(s.Code()))
	fields[1] = zap.String("rpc_method", info.FullMethod)
	fields[2] = zap.ByteString("request_body", req)
	fields[3] = zap.String("duration", du.String())

	if s.Code() != 0 && s.Code() < 1000 {
		respB, _ := json.Marshal(resp)
		if err != nil {
			respB = []byte(err.Error())
		}
		fields[4] = zap.ByteString("response", respB)
		logger.For(ctx).Error("failed", fields...)
	} else {
		fields[4] = zap.ByteString("response", []byte(""))
		logger.For(ctx).Info("success", fields...)
	}
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
		rpcStdLogger(ctx, info, reqB, respB, s, err, start)
		return
	}
}
