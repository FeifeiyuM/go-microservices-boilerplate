
package conf

import (
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	Http *httpConf `mapstructure:"http"`
	Nsq *nsqConf `mapstructure:"nsq"`
	Rpc *grpcConf `mapstructure:"rpc"`
}

// http config
type httpConf struct {
	Address string `mapstructure:"address"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
}

// nsq config
type nsqConf struct {
	NsqLookupds []string `mapstructure:"nsq_lookupds"`
	PollInterval time.Duration `mapstructure:"poll_interval"`
	MaxInFlight int `mapstructure:"max_in_flight`
}

// grpc config
type grpcConf struct {
	RpcPort string `mapstructure:"rpc_port"`
}

func NewConfig() *Config {
	return &Config{
		Http: &httpConf{
			Address: "",
			ReadTimeout:       20 * time.Second,
			WriteTimeout:      20 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			IdleTimeout:       10 * time.Second,
		},
		Nsq: &nsqConf{
			NsqLookupds: []string{},
			PollInterval: 10 * time.Second,
			MaxInFlight: 1,
		},
		Rpc: &grpcConf{
			RpcPort: "",
		},
	}
}

func Init(cfgFile string) (*Config, error) {
	if cfgFile == "" {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	} else {
		viper.SetConfigFile(cfgFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := NewConfig()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}