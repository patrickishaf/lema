package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string  `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Posts    []Post  `json:"posts"`
}
