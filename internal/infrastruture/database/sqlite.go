package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(db *sql.DB) {
	// SQL statement for creating a new table
	createTableSQL := `CREATE TABLE IF NOT EXISTS task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		priority TEXT NOT NULL
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func CloseDB(db *sql.DB) {
	db.Close()
}
