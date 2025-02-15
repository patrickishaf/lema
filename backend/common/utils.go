package common

import (
	"crypto/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func CreateErrorResponse(message string) map[string]any {
	return map[string]any{
		"status":  "error",
		"message": message,
	}
}

func SendResponse(c *gin.Context, statusCode int, data any, message string) {
	response := Response{
		Status:  "error",
		Message: message,
		Data:    data,
	}

	if statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
		response.Status = "success"
	}
	c.IndentedJSON(statusCode, response)
}

type PaginatedItems[T any] struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
	Data       []T `json:"data"`
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
