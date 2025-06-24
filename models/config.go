package models

import "os"

type Config struct {
	Port string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &Config{
		Port: port,
	}
}
