package httpx

import (
	"net"
	"time"

	"app/pkg/core/cfg"
)

const (
	CfgKeyPort    cfg.Key = "HTTP_SERVER_PORT"
	CfgKeyBaseURL cfg.Key = "HTTP_SERVER_BASE_URL"

	CfgDefaultPort    = "3000"
	CfgDefaultBaseURL = "/api/"
)

type Config struct {
	Port              string
	ReadHeaderTimeout time.Duration
	Timeout           time.Duration
	CacheTTL          time.Duration
	BaseURL           string
	Rate              RateLimit
	CORS              CORSConfig
	FrontMinVersion   string
}

type RateLimit struct {
	MaxRate int
	Every   time.Duration
}

type CORSConfig struct {
	Origins     string
	Methods     string
	Headers     string
	Credentials bool
}

// Addr returns server address in format ":<port>".
func (c Config) Addr() string {
	return net.JoinHostPort("", c.Port)
}
