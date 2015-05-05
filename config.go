package main

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

func ConfigInit() Config {
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println(err)
	}

	return config
}
