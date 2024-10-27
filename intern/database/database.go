package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error

	// Connect to the database
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")
}

func InitDB() {
	// SQL statement for creating a new table
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		priority TEXT NOT NULL,
		created_by TEXT NOT NULL,
		created_date TEXT NOT NULL,
		updated_by TEXT NOT NULL,
		updated_date TEXT NOT NULL
    );`

	// Create tasks table
	statement, err := DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Task table created or already exists")
}
