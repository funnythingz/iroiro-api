package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	AccessKey string `toml:"access_key"`
}

func AppConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("config/config.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}

var (
	Params = AppConfig()
)
