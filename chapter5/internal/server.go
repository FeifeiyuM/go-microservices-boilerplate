package internal

import (
	"context"
	"fmt"
	"github.com/FeifeiyuM/sqly"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"gmb/internal/interface/rpc"
	"gmb/pkg/log"
	"gmb/pkg/middlewares"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"

	"github.com/labstack/echo/v4"

	"gmb/internal/conf"
	"gmb/internal/dao"
	mq_handler "gmb/internal/interface/mq"
	"gmb/internal/interface/rest"
	"gmb/internal/service"
	"gmb/pkg/mq"
)

type server struct {
	logger log.Factory
	cfg    *conf.Config
	echo   *echo.Echo
	nsq    *mq.NsqConsumer
	grpc   *grpc.Server
}

func NewServer(cfg *conf.Config, logger log.Factory) *server {
	return &server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *server) initRestHandler() {
	e := echo.New()
	s.echo = e
	// 中间件引入
	s.echo.Use(middlewares.EchoTracer, middlewares.EchoRequestID, middlewares.EchoStandardLogger,
		middlewares.EchoErrorHandler, middlewares.EchoRecover)
	// account
	accRouter := e.Group("/account")
	accHandler := rest.NewHttpHandler(s.logger)
	accHandler.Router(accRouter)

}

func (s *server) initNsqHandler() {
	s.nsq = mq.NewNsqConsumer(s.cfg.Nsq.NsqLookupds, s.cfg.Nsq.PollInterval, s.cfg.Nsq.MaxInFlight)
	// 中间件引入
	s.nsq.Use(middlewares.MqTracer, middlewares.MqStdLogger, middlewares.MqErrHandler, middlewares.MqRecovery)
	// register mq handler
	mqHandler := mq_handler.NewMqHandler(s.logger)
	mqHandler.Register(s.nsq)
}

func (s *server) initRpcHandler() {
	s.grpc = grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			middlewares.GrpcUnaryTracerInterceptor(),
			middlewares.GrpcUnaryRequestIDInterceptor(),
			middlewares.GrpcUnaryStdLoggerInterceptor(),
			middlewares.GrpcUnaryRecoveryInterceptor(),
			middlewares.GrpcUnaryValidatorInterceptor(),
			middlewares.GrpcUnaryErrorInterceptor(),
		)))
	// register rpc handler
	rpcHandler := rpc.NewRpcHandler(s.logger)
	rpcHandler.Register(s.grpc)
}

func (s *server) initDaoLayer() {
	dbOpt := &sqly.Option{}
	db, err := sqly.New(dbOpt)
	if err != nil {
		panic(err)
	}
	dao.InitDao(db, nil, s.logger)
	// init account dao
	dao.InitAccountRepo()
	// init order dao
	dao.InitOrderRepo()
	// init property dao
	dao.InitPropertyRepo()
}

func (s *server) initServiceLayer() {
	// init account service
	service.InitAccountSrv(s.cfg, s.logger)
	// init order service
	service.InitOrderSrv(s.cfg, s.logger)
}

func (s *server) Run() {
	middlewares.InitLogger(s.logger)
	// dao layer
	s.initDaoLayer()
	// service layer
	s.initServiceLayer()
	// rest handler layer
	s.initRestHandler()
	// mq handler layer
	s.initNsqHandler()
	// rpc handler layer
	s.initRpcHandler()
	// run server
	// start rest server
	go func() {
		if err := s.echo.Start(s.cfg.Http.Address); err != nil {
			s.logger.Bg().Panic("failed to start http server", zap.Error(err))
		}
	}()
	// start rpc server
	go func() {
		lis, err := net.Listen("tcp", s.cfg.Rpc.RpcPort)
		if err != nil {
			s.logger.Bg().Panic("failed to start rpc server", zap.Error(err))
		}
		fmt.Println("grpc started...")
		if err = s.grpc.Serve(lis); err != nil {
			s.logger.Bg().Panic("failed to start rpc server", zap.Error(err))
		}
	}()
	// start nsq
	if err := s.nsq.Start(); err != nil {
		s.logger.Bg().Panic("failed to start nsq consumer", zap.Error(err))
	}
	fmt.Println("nsq started...")
}

func (s *server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.echo != nil {
		_ = s.echo.Shutdown(ctx)
	}
	if s.nsq != nil {
		_ = s.nsq.Close
	}
	if s.grpc != nil {
		s.grpc.GracefulStop()
	}
}
