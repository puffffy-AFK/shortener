package app

import (
	"fmt"
	"net/http"
	"shortener/internal/client/db"
	"shortener/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func registerRoutes(router *gin.Engine, log *logger.Logger, dbConn *db.Database) {
	router.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("Invalid request", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		shortCode := generateShortCode(req.URL)

		if err := dbConn.SaveURL(shortCode, req.URL); err != nil {
			log.Error("Failed to save URL", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
			return
		}

		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	})

	router.GET("/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")

		originalURL, err := dbConn.GetURL(shortCode)
		if err != nil {
			log.Error("URL not found", zap.String("shortCode", shortCode))
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		c.Redirect(http.StatusFound, originalURL)
	})
}