package middlewares

import (
	"context"
	"gmb/pkg/gmberror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	"google.golang.org/grpc"
)


func GMBErrToGRPCErr(ge gmberror.GMBError) (error, bool) {
	switch ge.HttpStatus() {
	case http.StatusOK:
		return nil, false
	case http.StatusBadRequest:
		return status.Error(codes.InvalidArgument, ge.Error()), ge.SendMsg()
	case http.StatusRequestTimeout, http.StatusBadGateway:
		return status.Error(codes.DeadlineExceeded, ge.Error()), ge.SendMsg()
	case http.StatusNotFound:
		return status.Error(codes.NotFound, ge.Error()), ge.SendMsg()
	case http.StatusConflict:
		return status.Error(codes.AlreadyExists, ge.Error()), ge.SendMsg()
	case http.StatusForbidden:
		return status.Error(codes.PermissionDenied, ge.Error()), ge.SendMsg()
	case http.StatusInsufficientStorage:
		return status.Error(codes.ResourceExhausted, ge.Error()), ge.SendMsg()
	case http.StatusPreconditionFailed, http.StatusUnprocessableEntity, http.StatusPreconditionRequired:
		return status.Error(codes.FailedPrecondition, ge.Error()), ge.SendMsg()
	case http.StatusNotAcceptable:
		return status.Error(codes.Aborted, ge.Error()), ge.SendMsg()
	//case http.StatusRequestedRangeNotSatisfiable:
	//	return status.Error(codes.OutOfRange, ge.Error()), ge.SendMsg()
	case http.StatusNotImplemented:
		return status.Error(codes.Unimplemented, ge.Error()), ge.SendMsg()
	case http.StatusInternalServerError:
		return status.Error(codes.Internal, ge.Error()), ge.SendMsg()
	case http.StatusServiceUnavailable:
		return status.Error(codes.Unavailable, ge.Error()), ge.SendMsg()
	//case http.StatusGone:
	//	return status.Error(codes.DataLoss, ge.Error()), ge.SendMsg()
	default:
		return status.Error(codes.Unknown, ge.Error()), ge.SendMsg()
	}
}

// GrpcUnaryErrorInterceptor error handler
func GrpcUnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			_, ok := status.FromError(err)
			if ok {
				sendWarnMsg(err)
				return nil, err
			}
			ge, ok := err.(gmberror.GMBError)
			if !ok {
				sendWarnMsg(err)
				return nil, status.Error(codes.Unknown, err.Error())
			}
			err2, sendMsg := GMBErrToGRPCErr(ge)
			if sendMsg {
				sendWarnMsg(err2)
			}
			return nil, err2
		}
		return resp, nil
	}
}
