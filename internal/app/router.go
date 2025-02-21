package app

import (
	"fmt"
	"net/http"
	"shortener/internal/client/db"
	"shortener/internal/config"
	"shortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, log *logger.Logger, dbConn *db.Database, cfg *config.Config) {
	router.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("Invalid request", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		shortCode := generateShortCode(req.URL)

		shortURL := fmt.Sprintf("%s:%s/%s", cfg.Server.Host, cfg.Server.Port, shortCode)
		if err := dbConn.SaveURL(shortCode, req.URL); err != nil {
			log.Error("Failed to save URL", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")

		originalURL, err := dbConn.GetURL(shortCode)
		if err != nil {
			log.Error("Failed to retrieve URL", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		c.Redirect(http.StatusFound, originalURL)
	})
}