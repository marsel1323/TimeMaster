package service

import (
	"TimeMaster/internal/model"
	repository "TimeMaster/internal/repository/psql"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: taskRepo}
}

func (s *TaskService) Create(task *model.Task) error {
	// Implement code to create a task using the repository method
	return nil
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	// Implement code to get all tasks using the repository method
	return nil, nil
}

func (s *TaskService) GetById(id int) (*model.Task, error) {
	// Implement code to get a task by ID using the repository method
	return nil, nil
}
