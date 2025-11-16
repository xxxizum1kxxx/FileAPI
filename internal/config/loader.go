package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) (*Config, error) {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
