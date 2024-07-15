package app

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port     int    `yaml:"port" env-required:"true"`
	Postgres string `yaml:"postgres" env-required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	return &cfg, nil
}
