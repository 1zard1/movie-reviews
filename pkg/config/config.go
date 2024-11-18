package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpAddr string `default:"8080"`
	DBAddr   string
}

func NewConfig() (*Config, error) {
	var s Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
