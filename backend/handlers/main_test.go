package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrickishaf/lema/db"
	"github.com/patrickishaf/lema/models"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	routerV1 := router.Group("/v1")
	NewPostsHandler(db.NewPostsRepository()).RegisterRequestHandlers(routerV1)
	NewUsersHandler(db.NewUserRespository()).RegisterRequestHandlers(routerV1)
	NewHealthCheckHandler(db.NewHealthCheckRepository()).RegisterRequestHandlers(routerV1)

	return router
}

func setup() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("error loading .env.test file: %v", err)
	}

	db.InitializeDb()
}

func teardown() {
	migrator := db.GetTestDB().Migrator()
	migrator.DropTable(&models.User{})
	migrator.DropTable(&models.Post{})
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
