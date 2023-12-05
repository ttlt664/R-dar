package config

import (
	"github.com/spf13/viper"
)

const HttpPort = ":7777"

type Config struct {
	Log struct {
		Level     string `toml:"level"`
		Output    string `toml:"output"`
		Formatter string `toml:"formatter"`
		LogPath   string `toml:"log_path"`
	} `toml:"log"`
	
	OSS struct {
		Endpoint   string `toml:"endpoint"`
		AccessKey  string `toml:"access_key"`
		SecretKey  string `toml:"secret_key"`
		BucketName string `toml:"bucket"`
	} `toml:"oss"`
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
