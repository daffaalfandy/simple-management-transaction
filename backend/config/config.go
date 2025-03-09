package config

import (
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Database DatabaseConfig
}

func Init() *Config {
	return &Config{
		Database: initDBConfig(),
	}
}
