package internal

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"gmb/internal/interface/rest"
	"gmb/internal/conf"
	"gmb/internal/dao"
	"gmb/internal/service"
)

type server struct {
	cfg *conf.Config
	echo *echo.Echo
}

func NewServer(cfg *conf.Config) *server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) initRestHandler() {
	e := echo.New()
	s.echo = e

	// account
	accRouter := e.Group("/account")
	accHandler := rest.NewHttpHandler()
	accHandler.Router(accRouter)

}

func (s *server) initDaoLayer() {
	dbMock := dao.NewNewDataBase()
	dao.InitDao(dbMock)
	// init account dao 
	dao.InitAccountRepo()
}

func (s *server) initServiceLayer() {
	// init account service 
	service.InitAccountSrv()
}

func (s *server) Run() {
	// dao layer
	s.initDaoLayer()
	// service layer
	s.initServiceLayer()
	// handler layer
	s.initRestHandler()
	// run server
	go func() {
		if err := s.echo.Start(s.cfg.Http.Address); err != nil {
			panic(err)
		}
	}()
}

func (s *server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if s.echo != nil {
		_ = s.echo.Shutdown(ctx)
	}
}