package rpc

import (
	"context"
	"gmb/internal/proto/pb"
	"gmb/internal/service"
	"google.golang.org/grpc"
)

type rpcHandler struct {
	pb.UnimplementedAccountServer
}

func NewRpcHandler() *rpcHandler {
	return &rpcHandler{}
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