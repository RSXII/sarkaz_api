package database

import (
	"database/sql"
	"log"
	"os"
	"sarkaz_api/config"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
    Conn *sql.DB
}

var instance *Database

func InitDB() {
    config.LoadEnvironmentVariables()

    conn, err := sql.Open("mysql", os.Getenv("DSN"))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }

    instance = &Database{Conn: conn}
}

func GetDBInstance() *Database{
    return instance
}

func (db *Database) GetCharacterRepository() *CharacterRepository {
    return &CharacterRepository{db: db.Conn}
}
func (db *Database) GetUserRepository() *UserRepository {
    return &UserRepository{db: db.Conn}
}