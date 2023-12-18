package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   Server   `mapstructure:"server"`
	Postgres Postgres `mapstructure:"postgres"`
}

type Postgres struct {
	Url string `mapstructure:"url"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

func LoadAppConfig(path string) (Configuration, error) {
	if path == "" {
		return Configuration{}, fmt.Errorf("config path is empty")
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Configuration{}, fmt.Errorf("config file not found: %s", path)
		}
		return Configuration{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var configuration Configuration

	if err := viper.Unmarshal(&configuration); err != nil {
		return Configuration{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return configuration, nil
}
