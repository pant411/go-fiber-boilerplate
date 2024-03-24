package model

import (
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed"`
}

// Validate validates the Todo struct
func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	return nil
}
