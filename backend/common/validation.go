package common

import (
	"github.com/go-playground/validator/v10"
)

const (
	emailSchema = "required,email"
)

var validate *validator.Validate

func ValidateStruct(data any) map[string]string {
	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(data)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return errors
	}
	return nil
}

func ValidateVariable(myVar any, schema string) error {
	err := validate.Var(myVar, schema)

	if err != nil {
		return err
	}

	return nil
}
