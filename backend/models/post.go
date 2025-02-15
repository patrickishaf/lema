package models

type Post struct {
	ID     string `json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
