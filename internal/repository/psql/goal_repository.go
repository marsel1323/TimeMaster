package repository

import (
	"TimeMaster/internal/model"
	"database/sql"
)

type GoalRepository struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) *GoalRepository {
	return &GoalRepository{db: db}
}

func (r *GoalRepository) Create(goal *model.Goal) error {
	// Implement code to insert goal into the database
	return nil
}

func (r *GoalRepository) GetAll() ([]*model.Goal, error) {
	// Implement code to retrieve all goals from the database
	return nil, nil
}

func (r *GoalRepository) GetByID(id int) (*model.Goal, error) {
	// Implement code to retrieve a goal by ID from the database
	return nil, nil
}
