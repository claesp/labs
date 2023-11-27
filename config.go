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
		port_i, port_int_err := strconv.Atoi(os.Getenv("LABS_PORT"))
		if port_int_err == nil {
			cfg.Port = port_i
		}
	}

	return cfg
}
