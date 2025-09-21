package usecase

import (
	"app/internal/entity"

	"github.com/rs/zerolog/log"
)

func (u *UseCasesImpl) CreateArticle(req *entity.CreateArticle) (*entity.CreateArticleResult, error) {
	result, err := u.Storage.CreateArticle(*entity.MakeCreateArticleEntityToDB(req))

	if err != nil {
		log.Error().Err(err).Msg("failed to CreateArticle")
		return nil, err
	}

	return entity.MakeСreateArticleDBToEntity(*result), nil
}
