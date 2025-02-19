package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		DSN string `env:"DATABASE_DSN" env-default:"postgres://user:password@localhost:5432/shortener?sslmode=disable"`
	} `yaml:"database"`
}

func LoadConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	return &cfg
}