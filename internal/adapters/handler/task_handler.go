package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-homework/internal/entities"
	"go-homework/internal/usecases"
)

type TaskHandler struct {
	TaskUsecase usecases.TaskUsecase
}

func NewTaskHandler(taskUsecase usecases.TaskUsecase) (h TaskHandler) {
	return TaskHandler{TaskUsecase: taskUsecase}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTasks(w, r)
	case http.MethodPost:
		h.CreateTask(w, r)
	case http.MethodPut:
		h.UpdateTask(w, r)
	case http.MethodDelete:
		h.DeleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) TasksRoute(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", h.HandleTasks)
	mux.HandleFunc("/tasks/{id}", h.GetTaskByID)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := h.TaskUsecase.Get(strconv.Itoa(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	response, err := h.TaskUsecase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task entities.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := h.TaskUsecase.Create(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var task entities.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := h.TaskUsecase.Update(strconv.Itoa(id), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.TaskUsecase.Delete(strconv.Itoa(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
