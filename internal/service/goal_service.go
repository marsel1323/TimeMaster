package service

import (
	"TimeMaster/internal/model"
	repository "TimeMaster/internal/repository/psql"
)

type GoalService struct {
	repo *repository.GoalRepository
}

func NewGoalService(goalRepo *repository.GoalRepository) *GoalService {
	return &GoalService{repo: goalRepo}
}

func (s *GoalService) Create(goal *model.Goal) error {
	// Implement code to create a goal using the repository method
	return nil
}

func (s *GoalService) GetAll() ([]model.Goal, error) {
	// Implement code to get all goals using the repository method
	return nil, nil
}

func (s *GoalService) GetById(id int) (*model.Goal, error) {
	// Implement code to get a goal by ID using the repository method
	return nil, nil
}
