package model

import (
	"errors"
	"fmt"
)

var (
	IDNotFound                 int64
	ErrTitleRequired           = errors.New("title cannot be null")
	ErrActivityGroupIDRequired = errors.New("activity_group_id cannot be null")
	ErrActivityNotFound        = fmt.Errorf("Activity with ID %d not found", IDNotFound)
	ErrTodoNotFound            = fmt.Errorf("Todo with ID %d not found", IDNotFound)
	// ErrEmailRequired = errors.New("email cannot be null")
)
