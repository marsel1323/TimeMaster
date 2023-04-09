package repository

import (
	"TimeMaster/internal/model"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	// Implement code to insert category into the database
	return nil
}

func (r *CategoryRepository) GetAll() ([]*model.Category, error) {
	// Implement code to retrieve all categories from the database
	return nil, nil
}

func (r *CategoryRepository) GetByID(id int) (*model.Category, error) {
	// Implement code to retrieve a category by ID from the database
	return nil, nil
}
