package main

import (
	"log"
	"net/http"

	"go-homework/internal/adapters/handler"
	"go-homework/internal/adapters/repository"
	"go-homework/internal/infrastruture/database"
	"go-homework/internal/usecases"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to the database
	dbConn, err := database.NewSqliteDB("./tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	// Close the database connection when the program exits
	defer database.CloseDB(dbConn)

	// Initialize the database
	database.InitDB(dbConn)

	// Initialize the task repository and usecase
	taskRepository := repository.NewTaskRepository(dbConn)
	taskUsecase := usecases.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// Set up HTTP routes
	mux := http.NewServeMux()
	taskHandler.TasksRoute(mux)

	http.Handle("/", mux)

	// Start HTTP server
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
