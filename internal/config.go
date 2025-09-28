package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Addr     string           `yaml:"port"`
	Database DatabaseSettings `yaml:"database"`
}

type DatabaseSettings struct {
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"database_name"`
}

func (ds *DatabaseSettings) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		ds.Username, ds.Password, ds.Host, ds.Port, ds.DbName,
	)
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
