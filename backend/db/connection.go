package db

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/patrickishaf/lema/models"
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

	configureConnectionPooling(database)

	db = database
	dbError := db.AutoMigrate(&models.Post{}, &models.User{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}

	// seedDatabase()
}

func getDB() *gorm.DB {
	return db
}

func configureConnectionPooling(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get underliying database")
	}

	maxOpenConnections, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))
	if err != nil {
		panic("ENV Error: invalid MAX_OPEN_CONNECTIONS.")
	}

	maxIdleConnections, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	if err != nil {
		panic("ENV Error: invalid MAX_IDLE_CONNECTIONS.")
	}

	maxIdleTimeSeconds, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTION_TIME_SECONDS"))
	if err != nil {
		panic("ENV Error: invalid MAX_IDLE_CONNECTION_TIME_SECONDS.")
	}
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(maxIdleTimeSeconds))
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to fetch inner db. error: %v", err)
	}
	sqlDB.Close()
}
