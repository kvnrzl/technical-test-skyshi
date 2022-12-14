package mysql

import (
	"context"
	"strconv"
	"technical_test_skyshi/model"
	"technical_test_skyshi/todo/repository"

	"gorm.io/gorm"
)

type MysqlTodo struct{}

func NewMysqlTodo() repository.TodoRepository {
	return &MysqlTodo{}
}

func (m *MysqlTodo) Create(ctx context.Context, db *gorm.DB, todo *model.TodoCreate) (int64, error) {
	// do some convert here
	parseActivityGroupID := strconv.Itoa(int(todo.ActivityGroupID))

	var isActive string
	if todo.IsActive == true {
		isActive = "1"
	} else {
		isActive = "0"
	}

	data := &model.Todo{
		Title:           todo.Title,
		ActivityGroupID: parseActivityGroupID,
		IsActive:        isActive,
		Priority:        todo.Priority,
	}
	err := db.WithContext(ctx).Create(data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (m *MysqlTodo) GetAllWithQueryParams(ctx context.Context, db *gorm.DB, queryParams string) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := db.WithContext(ctx).Where("activity_group_id = ?", queryParams).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (m *MysqlTodo) GetAllWithoutQueryParams(ctx context.Context, db *gorm.DB) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := db.WithContext(ctx).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (m *MysqlTodo) GetByID(ctx context.Context, db *gorm.DB, id int64) (*model.Todo, error) {
	var todo model.Todo
	err := db.WithContext(ctx).Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (m *MysqlTodo) Update(ctx context.Context, db *gorm.DB, todo *model.Todo) (*model.Todo, error) {
	err := db.WithContext(ctx).Where("id = ?", todo.ID).Updates(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (m *MysqlTodo) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Where("id = ?", id).Delete(&model.Todo{}).Error
	if err != nil {
		return err
	}
	return nil
}
