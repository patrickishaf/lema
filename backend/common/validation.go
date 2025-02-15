package common

import (
	"github.com/go-playground/validator/v10"
)

const (
	SchemaEmail      = "required,email"
	SchemaPageNumber = "required,min=1"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(data any) map[string]string {
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

func ValidateID(id string) error {
	err := validate.Var(id, "alphanumeric")
	if err != nil {
		return err
	}
	return nil
}
