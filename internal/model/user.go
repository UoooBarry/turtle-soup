package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
}
