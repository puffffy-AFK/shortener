package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Server struct {
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"http://localhost"`
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		DSN string `yaml:"dsn" env:"DATABASE_DSN"`
	} `yaml:"database"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return &cfg
}