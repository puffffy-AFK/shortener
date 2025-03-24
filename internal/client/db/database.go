package db

import (
	"fmt"
	"errors"
	"gorm.io/driver/postgres"
	"shortener/internal/model"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Database{DB: db}, nil
}

func (d *Database) SaveURL(shortCode, originalURL string) error {
	return d.DB.Exec("INSERT INTO urls (short_code, original_url) VALUES (?, ?)", shortCode, originalURL).Error
}

func (d *Database) GetURL(shortCode string) (string, error) {
	var url model.URL
	result := d.DB.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("URL not found")
		}
		return "", fmt.Errorf("failed to retrieve URL: %w", result.Error)
	}
	return url.OriginalURL, nil
}