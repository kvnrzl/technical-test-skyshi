package repository

import (
	"context"
	"technical_test_skyshi/model"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	Create(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error)
	GetAll(ctx context.Context, db *gorm.DB) ([]*model.Activity, error)
	GetByID(ctx context.Context, db *gorm.DB, id int64) (*model.Activity, error)
	Update(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error)
	Delete(ctx context.Context, db *gorm.DB, id int64) error
}
