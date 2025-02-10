package common

func CreateErrorResponse(message string) map[string]any {
	return map[string]any{
		"status":  "error",
		"message": message,
	}
}
