package db

import (
	"log"
	"os"

	"github.com/patrickishaf/lema-be/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeDb() {
	dbName := os.Getenv("DB_NAME")
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}

	db = database
	dbError := db.AutoMigrate(&models.Post{}, &models.User{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}

	insertError := insertDummyUsers()
	if insertError != nil {
		log.Println("failed to indsert dummy users", insertError)
	}
	insertError2 := insertDummyPosts()
	if insertError2 != nil {
		log.Println("failed to indsert dummy posts", insertError2)
	}
}

func getDB() *gorm.DB {
	return db
}
