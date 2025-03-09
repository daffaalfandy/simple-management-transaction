package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var DB *sql.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func initDBConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
}

func (c *Config) ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name,
	)

	var err error
	maxRetries := 10
	for i := range make(([]int), maxRetries) {
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			fmt.Println("Connected to MySQL successfully!")
			return DB
		}

		fmt.Printf("Database not ready (attempt %d/%d), retrying in 5s...\n", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("Database is not responding: %v", err)
	return nil
}
