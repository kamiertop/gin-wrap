package config

import (
	"flag"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	configPath = flag.String("config", "config.toml", "config path")
	V          *viper.Viper
	Conf       config
)

type config struct {
}

func Init() error {
	V = viper.New()
	V.SetConfigFile(*configPath)
	if err := V.ReadInConfig(); err != nil {
		return fmt.Errorf("read config file failed, err:%w", err)
	}
	if err := V.Unmarshal(&Conf, func(cfg *mapstructure.DecoderConfig) {
		cfg.TagName = "toml"
	}); err != nil {
		return fmt.Errorf("unmarshal config file failed, err:%w", err)
	}
	fmt.Printf("init config success, config:%#v\n", Conf)

	return nil
}
