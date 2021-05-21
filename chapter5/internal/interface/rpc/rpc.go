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

func (h *rpcHandler) Register(s *grpc.Server) {
	pb.RegisterAccountServer(s, h)
}

// 账号注册
func (h *rpcHandler) RegisterAccount(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	if in == nil {
		return nil, nil
	}
	errG := service.GetAccountSrv().CreateAccount(ctx, in.Name, in.Avatar, in.Mobile, int8(in.Gender))
	if errG != nil {
		return nil, errG
	}
	return &pb.RegisterResp{
		Msg: "ok",
	}, nil
}

func (h *rpcHandler) AccountRecharge(ctx context.Context, in *pb.AccountRechargeReq) (*pb.AccountRechargeReply, error) {
	if in == nil {
		return nil, nil
	}
	errG := service.GetOrderSrv().AccountRecharge(ctx, in.AccId, int64(in.Amount), in.PayOrderId)
	if errG != nil {
		return nil, errG
	}
	return &pb.AccountRechargeReply{
		Message: "OK",
	}, nil
}
