package cmd

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ConsumerKey    string `env:"TWITTER_CONSUMER_KEY,required"`
	ConsumerSecret string `env:"TWITTER_CONSUMER_SECRET,required"`
}

func ParseConfig() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return cfg, err
}
