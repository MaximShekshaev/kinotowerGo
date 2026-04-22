package core_database

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     int    `envconfig:"PORT" required:"true"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Name     string `envconfig:"DB" required:"true"`
	Database string `envconfig:"DB" required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("POSTGRES", &cfg) ; err != nil {
	return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}
	return &cfg, nil
}

func NewConfigMust() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(fmt.Errorf("cannot create config: %w", err))
	}
	return cfg
}