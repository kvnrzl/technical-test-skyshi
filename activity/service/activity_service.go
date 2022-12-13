package service

import (
	"context"
	"technical_test_skyshi/model"
)

type ActivityService interface {
	Create(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	GetAll(ctx context.Context) ([]*model.Activity, error)
	GetByID(ctx context.Context, id int64) (*model.Activity, error)
	Update(ctx context.Context, activity *model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, id int64) error
}
