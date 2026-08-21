package core

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	MongoURI               string
	MongoDatabase          string
	MongoProductCollection string
	ConnectionTimeout      time.Duration

	BffSvcPort     string
	ProductSvcPort string
	ProductSvcAddr string
}

func LoadConfig() (*Config, error) {
	cfg := Config{
		MongoURI:               os.Getenv("MONGO_URI"),
		MongoDatabase:          os.Getenv("MONGO_DATABASE"),
		MongoProductCollection: os.Getenv("MONGO_PRODUCT_COLLECTION"),
		ConnectionTimeout:      5 * time.Second,

		BffSvcPort:     os.Getenv("BFF_SVC_PORT"),
		ProductSvcPort: os.Getenv("PRODUCT_SVC_PORT"),
		ProductSvcAddr: os.Getenv("PRODUCT_SVC_ADDR"),
	}

	if cfg.MongoURI == "" ||
		cfg.MongoDatabase == "" ||
		cfg.MongoProductCollection == "" ||
		cfg.BffSvcPort == "" ||
		cfg.ProductSvcPort == "" ||
		cfg.ProductSvcAddr == "" {
		return nil, errors.New("env credentials required.")
	}

	return &cfg, nil
}
