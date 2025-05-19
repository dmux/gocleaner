package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// LoadConfigFunc é um tipo de função para carregar a configuração
type LoadConfigFunc func(string) (*Config, error)

type Schedule struct {
	Enabled bool   `yaml:"enabled"`
	Cron    string `yaml:"cron"` // Pode usar expressão cron ou `@every 1h`
}

type Config struct {
	Directory     string `yaml:"directory"`
	DaysThreshold int    `yaml:"days_threshold"`
	SMTP          struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		To       string `yaml:"to"`
	} `yaml:"smtp"`
	Schedule Schedule `yaml:"schedule"`
}

var LoadConfig LoadConfigFunc = func(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
