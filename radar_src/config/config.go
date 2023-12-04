package config

import (
	"github.com/spf13/viper"
)

const HttpPort = ":7777"

type Config struct {
	Log struct {
		Level     string
		Output    string
		Formatter string
		LogPath   string
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
