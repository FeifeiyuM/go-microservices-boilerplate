package cmd

import (
	"flag"
	"github.com/uber/jaeger-lib/metrics"
	jexpvar "github.com/uber/jaeger-lib/metrics/expvar"
	"gmb/internal"
	"gmb/internal/conf"
	"gmb/pkg/log"
	middleware "gmb/pkg/middlewares"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

var (
	cfgFile        string
	metricsFactory metrics.Factory
	zapLogger      *zap.Logger
)

// initLogger logger initialize
func initLogger() {
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.ISO8601TimeEncoder
	c.TimeKey = "time"
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(c),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.InfoLevel))
	zapLogger = zap.New(core, zap.AddStacktrace(zapcore.FatalLevel), zap.AddCaller())
}

func init() {
	// 从启动命令中读取配置文件路径
	flag.StringVar(&cfgFile, "c", "./config.toml", "path of mall config file.")
	metricsFactory = jexpvar.NewFactory(10)
	// 日志初始化
	initLogger()
}

func StartServer() {
	logger := log.NewFactory(zapLogger)
	// 初始化配置文件
	cfg, err := conf.Init(cfgFile)
	if err != nil || cfg == nil {
		logger.Bg().Fatal("failed to load config file")
		return
	}
	// 链路追踪
	_, tracingClose := middleware.TracingInit("gmb", metricsFactory)
	defer func() {
		_ = tracingClose.Close()
	}()
	// 启动服务
	server := internal.NewServer(cfg, logger)
	server.Run()

	// 开启系统信号接收通道
	// 防止系统推出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		server.Close()
	case syscall.SIGHUP:
	default:
	}
}
