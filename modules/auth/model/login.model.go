package model

import (
	"errors"
	"go-fiber-boilerplate/helpers/validation"
)

// User struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the Todo struct
func (l *Login) Validate() error {
	if l.Password == "" {
		return errors.New("password is required")
	}

	if l.Email == "" {
		return errors.New("email is required")
	}

	if !validation.ValidationEmail(l.Email) {
		return errors.New("email not valid")
	}

	return nil
}
