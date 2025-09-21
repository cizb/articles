package config

import (
	"strings"

	"app/pkg/core/cfg/viperx"
	"app/pkg/db"
	"app/pkg/httpx"

	"github.com/spf13/viper"
)

const (
	AppPrefix = "APP"
	keyDebug  = "DEBUG"
)

type Application struct {
	Debug       bool
	Environment string
	AppPrefix   string
}

type (
	Config struct {
		App        *Application
		PostgresDB *db.Config
		HTTP       *httpx.Config
	}
)

func LoadFromEnv() *Config {
	loader := viperx.EnvLoader(AppPrefix)
	config := &Config{
		App:        applicationFromEnv(loader),
		PostgresDB: db.CfgFromViper(loader),
		HTTP:       httpx.CfgFromViper(loader),
	}
	return config
}

func applicationFromEnv(loader *viper.Viper) *Application {
	env := viperx.Environment(loader)

	return &Application{
		Debug:       loader.GetBool(keyDebug),
		Environment: env.String(),
		AppPrefix:   strings.ToLower(AppPrefix),
	}
}
