package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Street   string `json:"street"`
	State    string `json:"state"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
	Posts    []Post `json:"posts"`
}
