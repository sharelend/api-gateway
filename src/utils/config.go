package utils

import (
	"log"

	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services map[string]string `yaml:"services"`
}

func LoadConfig() Config {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}
	return config
}
