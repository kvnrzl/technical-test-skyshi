package model

import (
	"errors"
)

// func NewErrActivityNotFound(id int) error {
// 	return fmt.Errorf("Activity with ID %v Not Found", id)
// }

var (
	ErrTitleRequired = errors.New("title cannot be null")
	// ErrActivityNotFound = NewErrActivityNotFound
	// ErrEmailRequired = errors.New("email cannot be null")
)
