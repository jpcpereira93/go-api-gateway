package config

import "github.com/rs/zerolog"

// Config represents the structure of the TOML configuration file.
type Config struct {
	Auth   AuthConfig   `toml:"auth"`
	DogsApi		DogsApiConfig	`toml:"dogs_api"`
	Logging   LoggingConfig   `toml:"logging"`
	Secret   SecretConfig   `toml:"secret"`
	Server   ServerConfig   `toml:"server"`
}

// AuthConfig holds secret related configs.
type AuthConfig struct {
	CookieName string `toml:"cookie_name"`
}

// DogsApiConfig holds the Dogs API related stuff.
type DogsApiConfig struct {
	BaseUrl string `toml:"base_url"`
}

// LoggingConfig holds logger related configs.
type LoggingConfig struct {
	Level zerolog.Level `toml:"level"`
	Output string `toml:"output"`
}

// SecretConfig holds secret related configs.
type SecretConfig struct {
	Iv string `toml:"iv"`
	Key string    `toml:"key"`
}

// ServerConfig holds server-related settings.
type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}
