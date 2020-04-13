package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Http HttpConfig `mapstructure:"http"`
}

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func NewConfig(path string) *Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	conf := &Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Could not unmarshal config: %s", err))
	}
	return conf
}


var WireSet = wire.NewSet(
	NewConfig,
)