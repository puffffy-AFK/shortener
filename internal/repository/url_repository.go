package repository

import (
	"shortener/internal/model"

	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url model.URL) error {
	return r.db.Create(&url).Error
}

func (r *URLRepository) FindByShortCode(shortCode string) (model.URL, error) {
	var url model.URL
	err := r.db.Where("short_code = ?", shortCode).First(&url).Error
	return url, err
}