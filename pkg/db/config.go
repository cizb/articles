package db

import (
	"app/pkg/core/cfg"
)

const (
	DriverPostgres = "postgres" // Имя драйвера PostgresSQL.

	SSLModeDisable    = "disable"     // Отключение SSL для подключения.
	SSLModeRequire    = "require"     // Требование SSL без проверки сертификата.
	SSLModeVerifyCA   = "verify-ca"   // SSL с проверкой сертификата от доверенного ЦС.
	SSLModeVerifyFull = "verify-full" // SSL с полной проверкой сертификата и соответствия имени хоста.

	CfgKeyHost        cfg.Key = "DB_HOST"          // Хост СУБД.
	CfgKeyPort        cfg.Key = "DB_PORT"          // Порт СУБД.
	CfgKeySSLMode     cfg.Key = "DB_SSL_MODE"      // Режим SSL.
	CfgKeySSLCertPath cfg.Key = "DB_SSL_CERT_PATH" // Путь к сертификату для SSL.
	CfgKeyDriver      cfg.Key = "DB_DRIVER"        // Имя драйвера БД.
	CfgKeyDatabase    cfg.Key = "DB_DATABASE"      // Имя БД.
	CfgKeyDBUser      cfg.Key = "DB_USER"          // Пользователь БД.
	CfgKeyDBPassword  cfg.Key = "DB_PASSWORD"      // Пароль пользователя БД.
	CfgKeyDebug       cfg.Key = "DB_DEBUG"         // Флаг отладки (bool).
	CfgKeyOptions     cfg.Key = "DB_OPTIONS"       // Опции подключения к БД.

	CfgDefaultHost        = "127.0.0.1" // Хост СУБД по-умолчанию.
	CfgDefaultPort        = "5432"      // Порт хоста СУБД по-умолчанию.
	CfgDefaultDataBase    = "postgres"
	CfgDefaultUser        = "postgres"
	CfgDefaultPassword    = "postgres"
	CfgDefaultSSLCertPath = ""
	CfgDefaultDebug       = false
	CfgDefaultMigration   = ""
	CfgDefaultDriver      = DriverPostgres // Драйвер БД по-умолчанию.
	CfgDefaultSSLMode     = SSLModeDisable // Режим SSL по-умолчанию.
)

type (
	Config struct {
		Driver      string // Имя драйвера БД.
		Database    string // Имя БД.
		Migration   string // Имя источника скриптов миграции БД.
		Debug       bool   // Флаг отладки.
		Host        string // Хост СУБД.
		Port        string // Порт СУБД.
		User        string // Пользователь БД.
		Password    string // Пароль пользователя БД.
		SSLMode     string // Режим SSL.
		SSLCertPath string // Путь к сертификату для SSL.
	}
)
