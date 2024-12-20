package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      `yaml:"app"`
	Http     `yaml:"http"`
	Log      `yaml:"logger"`
	Postgres `yaml:"postgres"`
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type Http struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
}

type Postgres struct {
	Url         string `env-required:"true" env:"POSTGRES_URL" yaml:"url"`
	MaxPoolSize int    `env-required:"true" env:"POSTGRES_MAX_POOL_SIZE" yaml:"max_pool_size"`
}

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return cfg, nil
}
