package db

import (
	"fmt"
	"gorm.io/driver/postgres"
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
	var originalURL string
	err := d.DB.Raw("SELECT original_url FROM urls WHERE short_code = ?", shortCode).Scan(&originalURL).Error
	if err != nil {
		return "", err
	}
	return originalURL, nil
}