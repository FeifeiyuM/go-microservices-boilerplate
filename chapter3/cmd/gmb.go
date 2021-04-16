package cmd

import (
	"github.com/uber/jaeger-lib/metrics"
	middleware "gmb/pkg/middlewares"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	"os"
	"os/signal"
	"syscall"
	"flag"
	"gmb/internal/conf"
	"gmb/internal"
)

var (
	cfgFile string
	metricsFactory metrics.Factory
)


func init() {
	// 从启动命令中读取配置文件路径
	flag.StringVar(&cfgFile, "c", "./config.toml", "path of mall config file.")
	metricsFactory = jexpvar.NewFactory(10)
}

func StartServer() {
	// 初始化配置文件
	cfg, err := conf.Init(cfgFile)
	if err != nil {
		panic(err)
	}
	// 链路追踪
	_, tracingClose := middleware.TracingInit("gmb", metricsFactory)
	defer func() {
		_ = tracingClose.Close()
	}()
	// 启动服务
	server := internal.NewServer(cfg)
	server.Run()
	
	// 开启系统信号接收通道
	// 防止系统推出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <- c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		server.Close()
	case syscall.SIGHUP:
	default:
	}
}