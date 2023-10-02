package database

import (
	"database/sql"
	"log"
	"os"
	"sarkaz_api/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    config.LoadEnvironmentVariables()

    db, err := sql.Open("mysql", os.Getenv("DSN"))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }

    DB = db
}
