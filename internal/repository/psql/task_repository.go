package repository

import (
	"TimeMaster/internal/model"
	"database/sql"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	// Implement code to insert task into the database
	return nil
}

func (r *TaskRepository) GetAll() ([]*model.Task, error) {
	// Implement code to retrieve all tasks from the database
	return nil, nil
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {
	// Implement code to retrieve a task by ID from the database
	return nil, nil
}
