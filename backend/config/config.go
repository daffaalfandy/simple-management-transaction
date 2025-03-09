package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"user", "password", "mysql", "3306", "app_db",
	)

	var err error
	for range time.Tick(5 * time.Second) {
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			fmt.Println("Connected to MySQL successfully!")
			return
		}

		fmt.Println("Database not ready, retrying in 5s...")
	}

	log.Fatal("Database is not responding:", err)
}
