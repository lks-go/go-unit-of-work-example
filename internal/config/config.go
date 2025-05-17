package config

import "github.com/ilyakaznacheev/cleanenv"

func Parse() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Addr string `env:"LISTEN_ADDR" env-default:":8080"`
	DSN  string `env:"DSN" env-default:"postgresql://admin:admin@localhost:5432/uow?sslmode=disable"`
}
