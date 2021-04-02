
package conf

import (
	"time"
	"github.com/spf13/viper"
)

type Config struct {
	Http *httpConf `mapstructure:"http"`
}

type httpConf struct {
	Address string `mapstructure:"address"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout"`
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