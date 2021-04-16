package middlewares

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)



// GrpcUnaryErrorInterceptor error handler
func GrpcUnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			// TODO 待完善错误处理
			fmt.Println(fmt.Sprintf("mq error: %s", err.Error()))
		}
		return resp, nil
	}
}
