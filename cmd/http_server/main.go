package main

import (
    "shortener/internal/app"
    "shortener/internal/client/db"
    "shortener/internal/config"
    "shortener/pkg/logger"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"

    _ "shortener/docs"
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

    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    go func() {
        if err := r.Run(":8081"); err != nil {
            log.Fatal("Failed to start server", err)
        }
    }()

    app.Run(log, dbConn, cfg)
}