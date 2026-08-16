// ws/services/user-svc/config/config.go
package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	MongoURI            string
	MongoDatabase       string
	MongoCollection       string
	Port                string
	ShutdownTimeout     time.Duration
	MongoConnectTimeout time.Duration

	TLSCertFile string
	TLSKeyFile  string
}

func LoadConfig() (*Config, error) {
	cfg := Config{
		MongoURI:            os.Getenv("MONGO_URI"),
		MongoDatabase:       os.Getenv("MONGO_DATABASE"),
		MongoCollection:       os.Getenv("MONGO_USER_COLLECTION"),
		Port:                os.Getenv("USER_SVC_PORT"),
		ShutdownTimeout:     10 * time.Second,
		MongoConnectTimeout: 10 * time.Second,

	}
	if cfg.MongoURI == "" || cfg.MongoDatabase == "" || cfg.MongoCollection == "" || cfg.Port == "" {
		return nil, errors.New("error loading db credentials")
	}
	return &cfg, nil
}
