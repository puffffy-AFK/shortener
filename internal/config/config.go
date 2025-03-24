package config

import (
    "github.com/gin-gonic/gin"
    "github.com/ilyakaznacheev/cleanenv"
    "log"
    "net/http"
)
type Config struct {

    Server struct {

        Host string `yaml:"host" env:"SERVER_HOST" env-default:"http://localhost"`

        Port string `yaml:"port" env:"SERVER_PORT" env-default:"8082"`
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
// @Summary Получить конфигурацию сервера
// @Description Возвращает текущую конфигурацию сервера, включая хост, порт и DSN базы данных.
// @Tags config
// @Produce  json
// @Success 200 {object} Config "Успешный ответ с конфигурацией"
// @Failure 500 {object} model.ErrorResponse "Ошибка сервера"
// @Router /config [get]
func GetConfigHandler(c *gin.Context) {
    cfg := LoadConfig()
    c.JSON(http.StatusOK, cfg)
}