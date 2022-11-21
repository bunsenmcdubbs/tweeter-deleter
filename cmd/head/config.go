package main

import (
	"github.com/caarlos0/env/v6"
)

type appConfig struct {
	ConsumerKey    string `env:"TWITTER_CONSUMER_KEY,required"`
	ConsumerSecret string `env:"TWITTER_CONSUMER_SECRET,required"`
}

func parseConfig() (appConfig, error) {
	cfg := appConfig{}
	err := env.Parse(&cfg)
	return cfg, err
}
