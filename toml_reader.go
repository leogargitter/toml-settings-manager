package main

import (
	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
)

// Define the structure for the config file
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
	Features FeaturesConfig
	Email    EmailConfig
	Advanced AdvancedConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string `toml:"dbname"`
}

type LoggerConfig struct {
	Level    string
	Filepath string
}

type FeaturesConfig struct {
	EnableCaching bool `toml:"enable_caching"`
	MaxCacheSize  int  `toml:"max_cache_size"`
}

type EmailConfig struct {
	SMTPServer  string `toml:"smtp_server"`
	SMTPPort    int    `toml:"smtp_port"`
	Username    string
	Password    string
	FromAddress string   `toml:"from_address"`
	ToAddresses []string `toml:"to_addresses"`
}

type AdvancedConfig struct {
	RetryAttempts int `toml:"retry_attempts"`
	Timeout       int
}

func GetConfig(tomlPath string) (Config, error) {
	var config Config

	if _, err := toml.DecodeFile(tomlPath, &config); err != nil {
		log.Warn("Error loading config:", err)
		return config, err
	}

	return config, nil
}
