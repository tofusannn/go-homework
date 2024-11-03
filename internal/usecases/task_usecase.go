package usecases

import (
	"go-homework/internal/adapters/repository"
	"go-homework/internal/entities"
)

type TaskUsecase struct {
	TaskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) (u TaskUsecase) {
	return TaskUsecase{TaskRepository: taskRepository}
}

func (u *TaskUsecase) Create(task entities.Task) (response entities.TaskResponse, err error) {
	return u.TaskRepository.Create(task)
}

func (u *TaskUsecase) Get(id string) (response entities.TaskResponse, err error) {
	return u.TaskRepository.Get(id)
}

func (u *TaskUsecase) GetAll() (response []entities.TaskResponse, err error) {
	return u.TaskRepository.GetAll()
}

func (u *TaskUsecase) Update(id string, task entities.Task) (response entities.TaskResponse, err error) {
	return u.TaskRepository.Update(id, task)
}

func (u *TaskUsecase) Delete(id string) (err error) {
	return u.TaskRepository.Delete(id)
}
