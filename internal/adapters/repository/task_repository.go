package repository

import (
	"database/sql"
	"go-homework/internal/entities"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) (r TaskRepository) {
	return TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task entities.Task) (response entities.TaskResponse, err error) {
	stmt, err := r.DB.Prepare("INSERT INTO task (title, description, priority) VALUES (?, ?, ?)")
	if err != nil {
		return response, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.Title, task.Description, task.Priority)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *TaskRepository) Get(id string) (response entities.TaskResponse, err error) {
	stmt, err := r.DB.Prepare("SELECT id, title, description, priority FROM task WHERE id = ?")
	if err != nil {
		return response, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&response.ID, &response.Title, &response.Description, &response.Priority)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *TaskRepository) GetAll() (response []entities.TaskResponse, err error) {
	rows, err := r.DB.Query("SELECT id, title, description, priority FROM task")
	if err != nil {
		return response, err
	}
	defer rows.Close()
	for rows.Next() {
		var task entities.TaskResponse
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority)
		if err != nil {
			return response, err
		}
		response = append(response, task)
	}
	return response, nil
}

func (r *TaskRepository) Update(id string, task entities.Task) (response entities.TaskResponse, err error) {
	stmt, err := r.DB.Prepare("UPDATE task SET title = ?, description = ?, priority = ? WHERE id = ?")
	if err != nil {
		return response, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.Title, task.Description, task.Priority, id)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *TaskRepository) Delete(id string) (err error) {
	stmt, err := r.DB.Prepare("DELETE FROM task WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
