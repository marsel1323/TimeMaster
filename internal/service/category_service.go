package service

import (
	"TimeMaster/internal/model"
	"TimeMaster/internal/repository/psql"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: categoryRepo}
}

func (s *CategoryService) Create(category *model.Category) error {
	// Implement code to create a category using the repository method
	return nil
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	// Implement code to get all categories using the repository method
	return nil, nil
}

func (s *CategoryService) GetById(id int) (*model.Category, error) {
	// Implement code to get a category by ID using the repository method
	return nil, nil
}
