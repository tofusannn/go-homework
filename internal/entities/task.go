package entities

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
}

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}
