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
	"github.com/joho/godotenv"
	db "github.com/patrickishaf/lema/db"
	handlers "github.com/patrickishaf/lema/handlers"
)

// TODO: Print the value of c.Query and see if it is a map or struct that you can validate
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env variables")
	}
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
	handlers.RegisterUserHandlers(router)
	handlers.RegisterPostHandlers(router)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("THE PORT IS...")
	fmt.Println(port)
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		fmt.Println("server is running on :8080...")
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
