package models

import "time"

type Post struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
