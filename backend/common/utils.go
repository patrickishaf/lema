package common

import (
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
