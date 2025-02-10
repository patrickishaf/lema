package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/patrickishaf/lema-be/db"
	"github.com/patrickishaf/lema-be/handlers"
)

// TODO: Print the value of c.Query and see if it is a map or struct that you can validate
func main() {
	db.InitializeDb()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))
	handlers.RegisterUserHandlers(router)
	handlers.RegisterPostHandlers(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
