package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Config

func init() {
	InitConfig()
}

func InitConfig() {
	filePath := "config.yaml"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func GetConfig() *Config {
	if config == nil {
		InitConfig()
	}

	return config
}
