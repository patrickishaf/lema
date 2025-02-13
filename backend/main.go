package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
	"github.com/patrickishaf/lema/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env variables")
	}
	db.InitializeDb()
	defer db.CloseDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	logger := common.InitializeLogger()
	router.Use(common.LogRequests(logger))

	handlers.RegisterUserHandlers(router)
	handlers.RegisterPostHandlers(router)
	handlers.RegisterHealthCheckHandlers(router)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
	fmt.Println("server is running on :8080...")
}
