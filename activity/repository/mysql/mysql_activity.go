package mysql

import (
	"context"
	"technical_test_skyshi/activity/repository"
	"technical_test_skyshi/model"

	"gorm.io/gorm"
)

type MysqlActivity struct{}

func NewMysqlAcitivity() repository.ActivityRepository {
	return &MysqlActivity{}
}

func (m *MysqlActivity) Create(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error) {
	err := db.WithContext(ctx).Create(activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (m *MysqlActivity) GetAll(ctx context.Context, db *gorm.DB) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := db.WithContext(ctx).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (m *MysqlActivity) GetByID(ctx context.Context, db *gorm.DB, id int64) (*model.Activity, error) {
	var activity model.Activity
	err := db.WithContext(ctx).Where("id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (m *MysqlActivity) Update(ctx context.Context, db *gorm.DB, activity *model.Activity) (*model.Activity, error) {
	// err := db.WithContext(ctx).Model(&activity).Where("id = ?", activity.ID).Update("title", activity.Title).Error
	err := db.WithContext(ctx).Where("id = ?", activity.ID).Updates(&activity).Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (m *MysqlActivity) Delete(ctx context.Context, db *gorm.DB, id int64) error {
	err := db.WithContext(ctx).Where("id = ?", id).Delete(&model.Activity{}).Error
	if err != nil {
		return err
	}
	return nil
}
