package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRespository() *UserRepository {
	return &UserRepository{
		db: getDB(),
	}
}

func (r *UserRepository) FindAll(pageNumber int, pageSize int) (*[]models.User, error) {
	var users []models.User
	var count int64

	offset := (pageNumber - 1) * pageSize
	err := getDB().Limit(pageSize).Offset(offset).Order("id asc, email").Find(&users).Error
	if err != nil {
		log.Printf("error in UserRepository.FindAll: %v", err)
		return nil, err
	}
	if err := getDB().Model(&models.User{}).Count(&count).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) GetUserCount() (int64, error) {
	var count int64
	if err := getDB().Model(&models.User{}).Count(&count).Error; err != nil {
		log.Printf("error in UserRepository.GetUserCount: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	if err := getDB().Where(&models.User{ID: id}).First(&user).Error; err != nil {
		log.Printf("error in UserRepository.FindUserById: %v", err)
		return &models.User{}, err
	}
	return &user, nil
}
