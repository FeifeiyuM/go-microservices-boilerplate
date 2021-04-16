package middlewares

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcUnaryRecoveryInterceptor grpc recovery
func GrpcUnaryRecoveryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = status.Errorf(codes.Internal, "%v", r)
				fmt.Println(fmt.Sprintf("grpc recovery error: %s", err.Error()))
			}
		}()
		resp, err = handler(ctx, req)
		panicked = false
		return
	}
}
