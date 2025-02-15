package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var newDB *gorm.DB

func initNewDB() {
	database, err := gorm.Open(sqlite.Open("newfile.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	newDB = database
	dbError := newDB.AutoMigrate(&models.Post{}, &models.Address{}, &models.User{})

	if dbError != nil {
		log.Println("failed to migrate database", dbError)
	}
}

func batchCopyUsers() {
	var users *[]models.User
	err := getDB().Find(&users).Error
	if err != nil {
		log.Println("failed to fetch users from old DB")
	}
	err = newDB.Create(&users).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
	}
}

func batchCopyAddresses() {
	var addresses *[]models.Address
	err := getDB().Find(&addresses).Error
	if err != nil {
		log.Println("failed to fetch addresses from old DB")
	}
	err = newDB.Create(&addresses).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
	}
}

func batchCopyPosts() {
	var posts *[]models.Post
	err := getDB().Find(&posts).Error
	if err != nil {
		log.Println("failed to fetch posts from old DB")
	}
	err = newDB.Create(&posts).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
	}
}

func funnelDB() {
	initNewDB()
	batchCopyUsers()
	batchCopyAddresses()
	batchCopyPosts()
}
