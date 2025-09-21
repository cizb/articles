package db

import (
	"github.com/spf13/viper"

	"app/pkg/core/cfg"
	"app/pkg/core/cfg/viperx"
)

// CfgFromViper загружает и возвращает конфигурацию подключения к СУБД с помощью viper.
func CfgFromViper(loader *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	loader.SetDefault(CfgKeyDatabase.String(), CfgDefaultDataBase)
	loader.SetDefault(CfgKeyDBPassword.String(), CfgDefaultPassword)
	loader.SetDefault(CfgKeySSLCertPath.String(), CfgDefaultSSLCertPath)
	loader.SetDefault(CfgKeyDebug.String(), CfgDefaultDebug)

	return &Config{
		Driver:    viperx.Get(loader, CfgKeyDriver.Map(keyMapping...), CfgDefaultDriver),
		Host:      viperx.Get(loader, CfgKeyHost.Map(keyMapping...), CfgDefaultHost),
		Port:      viperx.Get(loader, CfgKeyPort.Map(keyMapping...), CfgDefaultPort),
		Database:  viperx.Get(loader, CfgKeyDatabase.Map(keyMapping...), CfgDefaultDataBase),
		User:      viperx.Get(loader, CfgKeyDBUser.Map(keyMapping...), CfgDefaultUser),
		Password:  viperx.Get(loader, CfgKeyDBPassword.Map(keyMapping...), CfgDefaultPassword),
		Debug:     viperx.Get(loader, CfgKeyDebug.Map(keyMapping...), CfgDefaultDebug),
	}
}
