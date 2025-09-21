package storage

import (
	"app/internal/storage/schema"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article schema.Article) (*schema.Article, error)
	GetArticleByID(id string) (*schema.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r *articleRepository) CreateArticle(article schema.Article) (*schema.Article, error) {
	var result schema.Article
	if err := r.db.Create(&article).Scan(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *articleRepository) GetArticleByID(id string) (*schema.Article, error) {
	var result schema.Article
	err := r.db.First(&result, id).Error
	return &result, err
}
