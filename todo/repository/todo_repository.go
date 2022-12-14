package repository

import (
	"context"
	"technical_test_skyshi/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, db *gorm.DB, todo *model.TodoCreate) (int64, error)
	GetAllWithQueryParams(ctx context.Context, db *gorm.DB, queryParams string) ([]*model.Todo, error)
	GetAllWithoutQueryParams(ctx context.Context, db *gorm.DB) ([]*model.Todo, error)
	GetByID(ctx context.Context, db *gorm.DB, id int64) (*model.Todo, error)
	Update(ctx context.Context, db *gorm.DB, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, db *gorm.DB, id int64) error
}
