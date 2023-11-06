package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type LocalConfig struct {
	Env    string `yaml:"env" env-default:"develop"`
	Server struct {
		Address      string        `yaml:"address"`
		ReadTimeout  time.Duration `yaml:"readTimeout" env-default:"3s"`
		WriteTimeout time.Duration `yaml:"writeTimeout" env-default:"3s"`
	} `yaml:""`
}

func NewLocalConfig() *LocalConfig {
	configPath := "config/local.yaml"

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg LocalConfig
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot reading config: %s", configPath)
	}

	return &cfg
}
