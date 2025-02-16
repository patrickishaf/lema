package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var oldDBName string = "old.db"
var oldDB *gorm.DB

func initOldDB() error {
	database, err := gorm.Open(sqlite.Open(oldDBName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	oldDB = database
	return nil
}

func batchCopyUsers() error {
	var users *[]models.User
	err := oldDB.Find(&users).Error
	if err != nil {
		log.Println("failed to fetch users from old DB")
		return err
	}
	err = getDB().Create(&users).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
		return err
	}

	return nil
}

func batchCopyAddresses() error {
	var addresses *[]models.Address
	err := oldDB.Find(&addresses).Error
	if err != nil {
		log.Println("failed to fetch addresses from old DB")
		return err
	}
	err = getDB().Create(&addresses).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
		return err
	}

	return nil
}

func batchCopyPosts() error {
	var posts *[]models.Post
	err := oldDB.Find(&posts).Error
	if err != nil {
		log.Println("failed to fetch posts from old DB")
		return err
	}
	err = getDB().Create(&posts).Error
	if err != nil {
		log.Printf("failed to copy posts into new db. error: %v", err)
		return err
	}

	return nil
}

func SeedFromOldDB() error {
	err := initOldDB()
	if err != nil {
		return err
	}
	err = batchCopyUsers()
	if err != nil {
		return err
	}
	err = batchCopyAddresses()
	if err != nil {
		return err
	}
	err = batchCopyPosts()
	if err != nil {
		return err
	}
	return nil
}
