package validation

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateData(data any) map[string]string {
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
