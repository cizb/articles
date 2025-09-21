package db

import (
	"app/internal/storage/schema"
	"fmt"
	"log"
	"os"
	"time"

	l "github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQL(cfg *Config) (db *gorm.DB, err error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host,     // Хост PostgreSQL
		cfg.Port,     // Порт PostgreSQL
		cfg.User,     // Имя пользователя
		cfg.Password, // Пароль
		cfg.Database, // Имя базы данных
	)

	// Конфигурация логгера
	pgLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Настройка вывода
		logger.Config{
			SlowThreshold:             time.Second, // Порог медленных запросов
			IgnoreRecordNotFoundError: true,        // Игнорировать ошибку "record not found"
			Colorful:                  true,        // Цветной вывод
		},
	)

	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: pgLogger, // Логируем запросы для отладки
	})

	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		db = db.Debug()
	}

	l.Info().Msg("Database connected")

	// Автосоздание / автообновление таблиц
	if err := db.AutoMigrate(&schema.Article{}); err != nil {
		l.Error().Err(err).Msg("Failed to AutoMigrate")
	} else {
		l.Info().Msg("AutoMigrate completed")
	}
	return
}
