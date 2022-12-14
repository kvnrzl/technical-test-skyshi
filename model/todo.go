package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              int64          `json:"id" gorm:"primaryKey"`
	Title           string         `json:"title"`
	ActivityGroupID string         `json:"activity_group_id"`
	IsActive        string         `json:"is_active"`
	Priority        string         `json:"priority"`
	CreatedAt       time.Time      `json:"created_at" gorm:"'DEFAULT:current_timestamp'"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"'DEFAULT:current_timestamp'"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TodoCreate struct {
	ID              int64     `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	ActivityGroupID int64     `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at" gorm:"'DEFAULT:current_timestamp'"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"'DEFAULT:current_timestamp'"`
}

func (t *Todo) TableName() string {
	return "todo"
}
