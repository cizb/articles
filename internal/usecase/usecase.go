package usecase

import (
	"app/internal/entity"
	"app/internal/storage"
)

type Article interface {
	CreateArticle(req *entity.CreateArticle) (*entity.CreateArticleResult, error)
	GetArticleByID(req *entity.GetArticleByID) (*entity.GetArticleByIDResult, error)
}

type UseCasesImpl struct {
	Article
	Storage storage.ArticleRepository
}

func NewArticeUsecase(r storage.ArticleRepository) *UseCasesImpl {
	return &UseCasesImpl{Storage: r}
}
