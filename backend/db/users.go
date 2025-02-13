package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
)

func FindUsers(limit int, offset int) ([]models.User, error) {
	var users []models.User
	err := getDB().Limit(limit).Offset(offset).Order("id asc, email").Find(&users).Error
	if err != nil {
		log.Printf("failed to find users. error: %v", err)
		return nil, err
	}
	return users, nil
}

func FindUserCount() int64 {
	var count int64
	getDB().Model(&models.User{}).Count(&count)
	return count
}

func FindUserById(id uint) models.User {
	var user models.User
	getDB().Where(&models.User{ID: id}).First(&user)
	return user
}
