package config

import (
	_ "github.com/go-sql-driver/mysql"
)

var cfg *Config

type Config struct {
	Database      DatabaseConfig
	Authenticator Authenticator
}

func Init() *Config {
	cfg = &Config{
		Database:      initDBConfig(),
		Authenticator: initAuthenticator(),
	}

	return cfg
}

func Get() *Config {
	return cfg
}
