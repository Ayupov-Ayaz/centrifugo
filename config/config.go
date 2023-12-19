package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"time"
)

const (
	defaultExpiration = 1 * time.Hour
	configFile        = "config.json"
)

//go:embed config.json
var configFileFs embed.FS

type Config struct {
	ApiKey     string        `json:"api_key"`
	Secret     string        `json:"token_hmac_secret_key"`
	Expiration time.Duration `json:"-"`
}

func New() (*Config, error) {
	data, err := configFileFs.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	if cfg.Expiration == 0 {
		cfg.Expiration = defaultExpiration
	}

	return cfg, nil
}
