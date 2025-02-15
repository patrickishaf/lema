package common

import (
	"crypto/rand"

	"github.com/go-playground/validator/v10"
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

func GenerateID() (string, error) {
	length := 32
	const alphasAndNumbers = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	result := make([]byte, length)
	for i, b := range bytes {
		result[i] = alphasAndNumbers[b%byte(len(alphasAndNumbers))]
	}

	return string(result), nil
}
