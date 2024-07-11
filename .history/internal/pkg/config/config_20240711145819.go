package config

import (
	"time"

	"github.com/brilianpmw/synapsis/presentation"
)

type Config struct {
	JWTSecretKey      string
	JWTExpirationTime time.Duration
}

var mainconfig *Config

func NewConfig() *Config {
	return &Config{
		JWTSecretKey:      string(presentation.JwtKey),
		JWTExpirationTime: 5 * time.Hour, // Default expiration time
	}
}

func GetJWTExpirationTime() time.Duration {
	if mainconfig == nil {
		return 0
	}
	return mainconfig.JWTExpirationTime
}
