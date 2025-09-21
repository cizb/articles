package usecase

import (
	"app/internal/entity"

	"github.com/rs/zerolog/log"
)

func (u *UseCasesImpl) GetArticleByID(req *entity.GetArticleByID) (*entity.GetArticleByIDResult, error) {
	result, err := u.Storage.GetArticleByID(req.ID)

	if err != nil {
		log.Error().Err(err).Msg("failed to GetIncomeIndividual")
		return nil, err
	}

	return entity.MakeGetArticleDBToEntity(*result), nil
}
