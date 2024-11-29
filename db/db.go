package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(dataSourceName string) *sql.DB {
	var DB *sql.DB
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return DB
}

func CreateTable(DB *sql.DB) {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        price REAL NOT NULL,
        created_at TEXT NOT NULL
    );`

	if _, err := DB.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
}
