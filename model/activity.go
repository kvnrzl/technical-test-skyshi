package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int64          `json:"id" gorm:"primaryKey"` // primary key: https://gorm.io/docs/models.html#Primary-Key
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at" gorm:"'DEFAULT:current_timestamp'"` // autoCreateTime: https://gorm.io/docs/create.html#Auto-Created-Timestamp
	UpdatedAt time.Time      `json:"updated_at" gorm:"'DEFAULT:current_timestamp'"` // autoUpdateTime: https://gorm.io/docs/create.html#Auto-Updated-Timestamp
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`                       // soft delete: https://gorm.io/docs/delete.html#Soft-Delete
}
