package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

// LoadConfig loads the configuration from a TOML file.
func LoadConfig(filepath string) (*Config, error) {
	// Open the .env TOML file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Parse the TOML file into the Config struct
	var config Config
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
