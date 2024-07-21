package app

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port       int           `yaml:"port" env-required:"true"`
	Postgres   string        `yaml:"postgres" env-required:"true"`
	SessionAge time.Duration `yaml:"session_age" env-required:"true"`
	ApiKey     string        `yaml:"api_key" env-required:"true"`
}

func NewConfig(cfgPath string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	return &cfg, nil
}
