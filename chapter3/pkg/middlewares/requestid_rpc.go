package middlewares

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GrpcUnaryRequestIDInterceptor grpc request id
func GrpcUnaryRequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}
		var requestID string
		if val, ok := md["x-request-id"]; ok {
			requestID = val[0]
		} else {
			uid, _ := uuid.NewUUID()
			requestID = uid.String()
		}
		ctx = context.WithValue(ctx, "_requestID", requestID)
		return handler(ctx, req)
	}
}
