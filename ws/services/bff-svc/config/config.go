// ws/services/bff-svc/config/config.go
package config

import (
	"errors"
	"os"
)

type Config struct {
	Port        string
	UserSvcAddr string
}

func LoadConfig() (*Config, error) {
	cfg := Config{
		Port:        os.Getenv("BFF_SVC_PORT"),
		UserSvcAddr: os.Getenv("USER_SVC_ADDR"),
	}
	if cfg.Port == "" || cfg.UserSvcAddr == "" {
		return nil, errors.New("env credentials are required")
	}
	return &cfg, nil
}
