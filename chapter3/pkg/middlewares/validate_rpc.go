package middlewares

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var validate = validator.New()

// GrpcUnaryValidatorInterceptor request param validator
func GrpcUnaryValidatorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err = validate.Struct(req); err != nil {
			err = status.Error(codes.InvalidArgument, err.Error())
			return
		}
		resp, err = handler(ctx, req)
		return
	}
}
