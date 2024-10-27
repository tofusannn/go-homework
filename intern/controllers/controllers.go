package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"go-homework/intern/database"
	"go-homework/intern/models"
)

var (
	tasks []models.TaskResponse
	task  models.TaskResponse
	res   models.Response
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		if id == "" {
			GetTasks(w, r)
		} else {
			GetTaskByID(w, r, id)
		}
	case "POST":
		createTask(w, r)
	case "PUT":
		updateTaskByID(w, r, id)
	case "DELETE":
		deleteTaskByID(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	priority := r.URL.Query().Get("priority")
	query := "SELECT * FROM tasks"
	args := []interface{}{}

	if priority != "" {
		query += " WHERE priority = ?"
		args = append(args, priority)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	tasks = []models.TaskResponse{}
	// Loop through the rows and add them to the tasks slice
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Created_By, &task.Created_Date, &task.Updated_By, &task.Updated_Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		res = models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Task not found",
			Data:       nil,
		}
	} else {

		res = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data:       tasks,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request, id string) {
	// Query the database for the task by ID
	row := database.DB.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	// Scan the row into the task variable
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Created_By, &task.Created_Date, &task.Updated_By, &task.Updated_Date)

	if err == sql.ErrNoRows {
		res = models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Task not found",
			Data:       nil,
		}
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		res = models.Response{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data:       task,
		}
	}

	json.NewEncoder(w).Encode(res)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest models.TaskRequest
	// Check if the body is empty
	if r.ContentLength == 0 {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	// Title and Priority cannot be empty
	if taskRequest.Title == "" || taskRequest.Priority == "" {
		http.Error(w, "Title and Priority are required", http.StatusBadRequest)
		return
	}

	// Priority must be low, medium or high
	if taskRequest.Priority != "low" && taskRequest.Priority != "medium" && taskRequest.Priority != "high" {
		http.Error(w, "Priority must be low, medium or high", http.StatusBadRequest)
		return
	}

	statement, err := database.DB.Prepare("INSERT INTO tasks(title, description, priority, created_by, created_date, updated_by, updated_date) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	Created_By := token
	Created_Date := time.Now().Format("2006-01-02 15:04:05")

	_, err = statement.Exec(taskRequest.Title, taskRequest.Description, taskRequest.Priority, Created_By, Created_Date, Created_By, Created_Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res = models.Response{
		StatusCode: http.StatusCreated,
		Message:    "Task created successfully",
		Data:       nil,
	}

	json.NewEncoder(w).Encode(res)
}

func updateTaskByID(w http.ResponseWriter, r *http.Request, id string) {
	var taskRequest models.TaskRequest
	// Check if the body is empty
	if r.ContentLength == 0 {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		http.Error(w, "Invalid request body format", http.StatusBadRequest)
		return
	}

	// Title and Priority cannot be empty
	if taskRequest.Title == "" || taskRequest.Priority == "" {
		http.Error(w, "Title and Priority are required", http.StatusBadRequest)
		return
	}

	// Priority must be low, medium or high
	if taskRequest.Priority != "low" && taskRequest.Priority != "medium" && taskRequest.Priority != "high" {
		http.Error(w, "Priority must be low, medium or high", http.StatusBadRequest)
		return
	}

	statement, err := database.DB.Prepare("UPDATE tasks SET title = ?, description = ?, priority = ?, updated_by = ?, updated_date = ? WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("Authorization")
	Updated_By := token
	Updated_Date := time.Now().Format("2006-01-02 15:04:05")

	_, err = statement.Exec(taskRequest.Title, taskRequest.Description, taskRequest.Priority, Updated_By, Updated_Date, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res = models.Response{
		StatusCode: http.StatusOK,
		Message:    "Task updated successfully",
		Data:       nil,
	}

	json.NewEncoder(w).Encode(res)
}

func deleteTaskByID(w http.ResponseWriter, id string) {
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		res := models.Response{
			StatusCode: http.StatusNotFound,
			Message:    "Task not found",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res = models.Response{
		StatusCode: http.StatusOK,
		Message:    "Task deleted successfully",
		Data:       nil,
	}

	json.NewEncoder(w).Encode(res)
}
