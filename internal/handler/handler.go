package handler

import (
	"app/config"
	"app/internal/usecase"
)

type Handler struct {
	useCase usecase.UseCasesImpl
	Config  *config.Config
}

func NewHandler(u *usecase.UseCasesImpl, cfg *config.Config) *Handler {
	return &Handler{
		useCase: *u,
		Config:  cfg,
	}
}
