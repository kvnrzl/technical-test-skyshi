package service

import (
	"context"
	"strconv"
	"technical_test_skyshi/model"
	"technical_test_skyshi/todo/repository"

	"gorm.io/gorm"
)

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
	db             *gorm.DB
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository, db *gorm.DB) TodoService {
	return &TodoServiceImpl{
		todoRepository: todoRepository,
		db:             db,
	}
}

func (s *TodoServiceImpl) Create(ctx context.Context, todo *model.TodoCreate) (*model.TodoCreate, error) {
	if todo.Title == "" {
		return nil, model.ErrTitleRequired
	}

	if todo.ActivityGroupID == 0 {
		return nil, model.ErrActivityGroupIDRequired
	}

	todo.IsActive = true
	todo.Priority = "very-high"

	id, err := s.todoRepository.Create(ctx, s.db, todo)
	if err != nil {
		return nil, err
	}

	newTodo, err := s.todoRepository.GetByID(ctx, s.db, id)
	if err != nil {
		return nil, err
	}

	parseActivityGroupID, err := strconv.Atoi(newTodo.ActivityGroupID)
	if err != nil {
		return nil, err
	}

	parseIsActive, err := strconv.ParseBool(newTodo.IsActive)
	if err != nil {
		return nil, err
	}

	return &model.TodoCreate{
		ID:              newTodo.ID,
		Title:           newTodo.Title,
		ActivityGroupID: int64(parseActivityGroupID),
		IsActive:        parseIsActive,
		Priority:        newTodo.Priority,
		CreatedAt:       newTodo.CreatedAt,
		UpdatedAt:       newTodo.UpdatedAt,
	}, nil
}

func (s *TodoServiceImpl) GetAll(ctx context.Context, queryParams string) ([]*model.Todo, error) {
	if queryParams == "" {
		return s.todoRepository.GetAllWithoutQueryParams(ctx, s.db)
	} else {
		return s.todoRepository.GetAllWithQueryParams(ctx, s.db, queryParams)
	}
}

func (s *TodoServiceImpl) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	res, err := s.todoRepository.GetByID(ctx, s.db, id)
	if err == gorm.ErrRecordNotFound {
		model.IDNotFound = id
		return nil, model.ErrTodoNotFound
	}

	return res, err
}

func (s *TodoServiceImpl) Update(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	_, err := s.GetByID(ctx, todo.ID)
	if err != nil {
		return nil, err
	}

	res, err := s.todoRepository.Update(ctx, s.db, todo)
	if err != nil {
		return nil, err
	}

	return s.todoRepository.GetByID(ctx, s.db, res.ID)
}

func (s *TodoServiceImpl) Delete(ctx context.Context, id int64) error {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.todoRepository.Delete(ctx, s.db, id)
}
