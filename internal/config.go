package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Addr string `yaml:"port"`
}

func GetConfig() (*Settings, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	var settings Settings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	// Override with env vars if present
	if val := os.Getenv("APP_PORT"); val != "" {
		settings.Addr = val
	}

	return &settings, nil
}
