package intern

import (
	"fmt"
	"go-homework/intern/controllers"
	"go-homework/intern/database"
	"go-homework/intern/middleware"
	"log"
	"net/http"
)

func Main() {
	fmt.Println("Simple REST API for a To-Do List App")

	// Connect to the database
	database.Connect()
	database.InitDB()

	setupRoutes()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.Handle("/tasks", middleware.TokenAuthMiddleware(http.HandlerFunc(controllers.HandleTasks)))
}
