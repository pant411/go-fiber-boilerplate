package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(field interface{}) error {
	// Create a new validator instance
	validate := validator.New()

	// Validate the struct
	if err := validate.Struct(field); err != nil {
		// Return validation error
		return err
	}

	return nil
}
