package config

import (
    "os"
)

type Config struct {
    DBUser string
    DBPass string
    DBHost string
    DBName string
}

func LoadConfig() *Config {
    return &Config{
        DBUser: os.Getenv("DB_USER"),
        DBPass: os.Getenv("DB_PASS"),
        DBHost: os.Getenv("DB_HOST"),
        DBName: os.Getenv("DB_NAME"),
    }
}
