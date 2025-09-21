package httpx

import (
	"github.com/spf13/viper"

	"app/pkg/core/cfg"
	"app/pkg/core/cfg/viperx"
)

func CfgFromViper(v *viper.Viper, keyMapping ...cfg.KeyMap) *Config {
	return &Config{
		Port:    viperx.Get(v, CfgKeyPort.Map(keyMapping...), CfgDefaultPort),
		BaseURL: viperx.Get(v, CfgKeyBaseURL.Map(keyMapping...), CfgDefaultBaseURL),
	}
}
