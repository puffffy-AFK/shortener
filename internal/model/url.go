package model

type URL struct {
	ID         uint   `gorm:"primaryKey"`
	ShortCode  string `gorm:"uniqueIndex"`
	OriginalURL string
}