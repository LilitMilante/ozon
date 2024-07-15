package app

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int `yaml:"Port" env-required:"true"`

	Database DBConfig `yaml:"Database"`
}

type DBConfig struct {
	Host     string `yaml:"DB_Host" env-required:"true"`
	Port     int    `yaml:"DB_Port" env-required:"true"`
	Name     string `yaml:"DB_Name" env-required:"true"`
	User     string `yaml:"DB_User" env-required:"true"`
	Password string `yaml:"DB_Password" env-required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg, nil
}

func validateConfig(config Config) error {
	if config.Port <= 0 {
		return fmt.Errorf("server port field is required")
	}
	if config.Database.Host == "" {
		return fmt.Errorf("database host field is required")
	}
	if config.Database.Port <= 0 {
		return fmt.Errorf("database port field is required")
	}
	if config.Database.Name == "" {
		return fmt.Errorf("database name field is required")
	}
	if config.Database.User == "" {
		return fmt.Errorf("database user field is required")
	}
	if config.Database.Password == "" {
		return fmt.Errorf("database password field is required")
	}

	return nil
}
