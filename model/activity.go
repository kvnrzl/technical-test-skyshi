package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at" gorm:"'DEFAULT:current_timestamp'"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"'DEFAULT:current_timestamp'"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type ActivityCreate struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at" gorm:"'DEFAULT:current_timestamp'"`
	UpdatedAt time.Time `json:"updated_at" gorm:"'DEFAULT:current_timestamp'"`
}

func (a *Activity) TableName() string {
	return "activities"
}
