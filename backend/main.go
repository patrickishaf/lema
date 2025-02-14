package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema/common"
	"github.com/patrickishaf/lema/db"
	"github.com/patrickishaf/lema/handlers"
)

func main() {
	db.InitializeDb()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))
	router.Use(common.PlaceRateLimits())

	logger := common.InitializeLogger()
	router.Use(common.LogRequests(logger))

	routerV1 := router.Group("/v1")

	handlers.NewPostsHandler(db.NewPostsRepository()).RegisterRequestHandlers(routerV1)
	handlers.NewUsersHandler(db.NewUserRespository()).RegisterRequestHandlers(routerV1)
	handlers.NewHealthCheckHandler(db.NewHealthCheckRepository()).RegisterRequestHandlers(routerV1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	addr := "0.0.0.0:" + port
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		fmt.Printf("server is running on %s...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("\nshutting down the server...")

	db.CloseDB()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %v", err)
	}
	fmt.Println("server has been shut down.")
}
