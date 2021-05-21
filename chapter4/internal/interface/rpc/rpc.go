package rpc

import (
	"context"
	"gmb/internal/proto/pb"
	"gmb/internal/service"
	"gmb/pkg/log"
	"google.golang.org/grpc"
)

type rpcHandler struct {
	logger log.Factory
	pb.UnimplementedAccountServer
}

func NewRpcHandler(logger log.Factory) *rpcHandler {
	return &rpcHandler{
		logger: logger,
	}
}

// 账号注册
func (h *rpcHandler) RegisterAccount(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	if req == nil {
		return nil, nil
	}

	err := service.GetAccountSrv().CreateAccount(ctx, req.Name, req.Address, int8(req.Gender))
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		Msg:                  "ok",
	}, nil
}

func (h *rpcHandler) Register(s *grpc.Server) {
	pb.RegisterAccountServer(s, h)
}