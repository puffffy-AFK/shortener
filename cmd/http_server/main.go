package main

import (
	"shortener/internal/app"
	"shortener/internal/client/db"
	"shortener/internal/config"
	"shortener/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()

	log := logger.NewLogger()

	dbConn, err := db.NewDatabase(cfg.Database.DSN)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	sqlDB, err := dbConn.DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB", err)
	}

	defer sqlDB.Close()

	app.Run(log, dbConn, cfg)
}