package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `validate:"required_on_create,optional_on_update" json:"title"`
	Completed bool   `json:"completed"`
}
