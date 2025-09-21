package main

import (
	"app/config"
	"app/internal/handler"
	"app/internal/router"
	"app/internal/storage"
	"app/internal/usecase"
	"app/pkg/httpx"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"app/pkg/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := initConfigs()

	r := gin.Default()

	if err := run(ctx, cfg, r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start application")
	}
}

func run(ctx context.Context, cfg *config.Config, r *gin.Engine) error {
	db, err := db.NewPostgreSQL(cfg.PostgresDB)

	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}

	repo := storage.NewArticleRepository(db)
	usecase := usecase.NewArticeUsecase(repo)
	handler := handler.NewHandler(usecase, cfg)

	router.RegisterRoutes(r, handler)
	err = httpx.StartServer(ctx, httpx.NewServer(cfg.HTTP, r))

	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func initConfigs() *config.Config {
	cfg := config.LoadFromEnv()
	return cfg
}
