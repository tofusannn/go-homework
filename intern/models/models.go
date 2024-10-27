package models

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type TaskRequest struct {
	Task
}

type TaskResponse struct {
	ID string `json:"id"`
	Task
	Created_By   string `json:"created_by"`
	Created_Date string `json:"created_date"`
	Updated_By   string `json:"updated_by"`
	Updated_Date string `json:"updated_date"`
}

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}
