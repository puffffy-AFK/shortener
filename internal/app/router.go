package app

import (
	"fmt"
	"net/http"
	"shortener/internal/client/db"
	"shortener/internal/config"
	"shortener/pkg/logger"
	"shortener/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary Создать короткую ссылку
// @Description Преобразует длинный URL в короткий код и сохраняет его в базе данных.
// @Tags links
// @Accept  json
// @Produce  json
// @Param   input body model.CreateShortURLRequest true "Длинный URL для сокращения"
// @Success 200 {object} model.CreateShortURLResponse "Успешный ответ с короткой ссылкой"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос"
// @Failure 500 {object} model.ErrorResponse "Ошибка сервера"
// @Router /shorten [post]
func shortenURLHandler(c *gin.Context, log *logger.Logger, dbConn *db.Database, cfg *config.Config) {
	var req model.CreateShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request"})
		return
	}

	shortCode := generateShortCode(req.URL)

	shortURL := fmt.Sprintf("%s:%s/%s", cfg.Server.Host, cfg.Server.Port, shortCode)
	if err := dbConn.SaveURL(shortCode, req.URL); err != nil {
		log.Error("Failed to save URL", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "Failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, model.CreateShortURLResponse{ShortURL: shortURL})
}

// @Summary Перенаправить по короткой ссылке
// @Description Перенаправляет пользователя на оригинальный URL, соответствующий короткому коду.
// @Tags links
// @Param   shortCode path string true "Короткий код"
// @Success 302 "Перенаправление на оригинальный URL"
// @Failure 404 {object} model.ErrorResponse "URL не найден"
// @Router /{shortCode} [get]
func redirectURLHandler(c *gin.Context, log *logger.Logger, dbConn *db.Database) {
	shortCode := c.Param("shortCode")

	originalURL, err := dbConn.GetURL(shortCode)
	if err != nil {
		log.Error("Failed to retrieve URL", err)
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}


func registerRoutes(router *gin.Engine, log *logger.Logger, dbConn *db.Database, cfg *config.Config) {
	router.POST("/shorten", func(c *gin.Context) {
		shortenURLHandler(c, log, dbConn, cfg)
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		redirectURLHandler(c, log, dbConn)
	})
}