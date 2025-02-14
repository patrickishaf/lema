package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	writer := makeRequest("GET", "/v1/health", "")
	assert.Equal(t, http.StatusOK, writer.Code)
}
