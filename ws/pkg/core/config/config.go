// /workspaces/gocline/ws/pkg/core/config/config.go
package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	MongoURI            string
	MongoDatabase       string
	MongoUserCollection string
	ConnectionTimeOut   time.Duration
	UserSvcPort         string
	UserSvcAddr         string
	BffSvcPort          string
}

func LoadConfig() (*Config, error) {
	cfg := Config{
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDatabase:       os.Getenv("MONGO_DB"),
		MongoUserCollection: os.Getenv("MONGO_USER_COLLECTION"),
		ConnectionTimeOut:   10 * time.Second,
		UserSvcPort:         os.Getenv("USER_SVC_PORT"),
		UserSvcAddr:         os.Getenv("USER_SVC_ADDR"),
		BffSvcPort:          os.Getenv("BFF_SVC_PORT"),
	}
	if cfg.MongoURI == "" ||
		cfg.MongoDatabase == "" ||
		cfg.MongoUserCollection == "" ||
		cfg.UserSvcPort == "" ||
		cfg.UserSvcAddr == "" ||
		cfg.BffSvcPort == "" {
		return nil, fmt.Errorf("env credentials are required.")
	}

	return &cfg, nil
}
