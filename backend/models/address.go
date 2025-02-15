package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID      string `json:"id" gorm:"primaryKey"`
	UserID  string `json:"user_id"`
	Street  string `json:"street"`
	State   string `json:"state"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}
