package model

//@description Эта структура используется для хранения данных в базе данных. Она содержит уникальный идентификатор, короткий код и оригинальный URL.
//@Property ID int `json:"id"` The unique ID of the URL record
//@Property ShortCode string `json:"short_code"` The short code for the URL
//@Property OriginalURL string `json:"original_url"` The original URL
type URL struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	ShortCode   string `gorm:"uniqueIndex" json:"short_code" example:"abc123"`
	OriginalURL string `json:"original_url" example:"https://example.com/long-url"`
}

// @description Запрос содержит оригинальный URL, который нужно сократить.
type CreateShortURLRequest struct {
	URL string `json:"url" example:"https://example.com/long-url"`
}

// @description Ответ содержит сокращенный URL.
type CreateShortURLResponse struct {
	ShortURL string `json:"short_url" example:"http://localhost:8080/abc123"`
}

// @description Ответ содержит сообщение об ошибке.
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}