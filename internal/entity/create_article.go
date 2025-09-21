package entity

import (
	"app/internal/storage/schema"
	"time"
)

type (
	CreateArticle struct {
		Title   string   `json:"title" validate:"required,min=3,max=255"`
		Tags    []string `json:"tags"`
		Content string   `json:"content" validate:"required"`
	}

	CreateArticleResult struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Tags      []string  `json:"tags"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func MakeCreateArticleEntityToDB(req *CreateArticle) *schema.Article {
	return &schema.Article{
		Title:   req.Title,
		Tags:    req.Tags,
		Content: req.Content,
	}
}

func MakeСreateArticleDBToEntity(req schema.Article) *CreateArticleResult {
	return &CreateArticleResult{
		ID:        req.ID,
		Title:     req.Title,
		Content:   req.Content,
		Tags:      req.Tags,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}
