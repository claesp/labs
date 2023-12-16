package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port int `json:"port"`
}

func loadConfigDefaults() Config {
	config := Config{}
	config.Port = 8080

	return config
}

func loadConfig() Config {
	cfg := loadConfigDefaults()

	if os.Getenv("LABS_PORT") != "" {
		portInt, portIntErr := strconv.Atoi(os.Getenv("LABS_PORT"))
		if portIntErr == nil {
			cfg.Port = portInt
		}
	}

	return cfg
}
