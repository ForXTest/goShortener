package main

import (
	"github.com/spf13/viper"
	"urlShotener/cmd/database"
)

type GinMode struct {
	Mode string
}

type ServerConfig struct {
	Host string
	Port int
}

type Config struct {
	Server  ServerConfig
	GinMode	GinMode
	MySQL	database.MySQLConfig
}

func GetConfig() (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
