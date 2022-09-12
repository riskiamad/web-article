package model

import "time"

// ArticleCreateRequest: struct hold body request for create article
type ArticleCreateRequest struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

// ArticleResponse: struct hold response format for article
type ArticleResponse struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
