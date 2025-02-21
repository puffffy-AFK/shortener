package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shortener/internal/client/db"
	"shortener/internal/config"
	"shortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run(log *logger.Logger, dbConn *db.Database, cfg *config.Config) {

	router := gin.Default()

	registerRoutes(router, log, dbConn, cfg)
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info("Starting server on port " + cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error", err)
		}
	}()

	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", err)
	}

	log.Info("Server exited properly")
}