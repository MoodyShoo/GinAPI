package database

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func LoadConfig() (*ConfigDatabase, error) {
	cfg := &ConfigDatabase{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("ошибка при загрузке конфигурации: %w", err)
	}

	return cfg, nil
}
