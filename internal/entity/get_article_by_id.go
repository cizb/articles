package entity

import (
	"app/internal/storage/schema"
	"time"
)

type (
	GetArticleByID struct {
		ID string `json:"id"`
	}

	GetArticleByIDResult struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Tags      []string  `json:"tags"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func MakeGetArticleDBToEntity(req schema.Article) *GetArticleByIDResult {
	return &GetArticleByIDResult{
		ID:        req.ID,
		Title:     req.Title,
		Content:   req.Content,
		Tags:      req.Tags,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}
