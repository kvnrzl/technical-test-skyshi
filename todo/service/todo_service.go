package service

import (
	"context"
	"technical_test_skyshi/model"
)

type TodoService interface {
	Create(ctx context.Context, todo *model.TodoCreate) (*model.TodoCreate, error)
	GetAll(ctx context.Context, queryParams string) ([]*model.Todo, error)
	GetByID(ctx context.Context, id int64) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, id int64) error
}
