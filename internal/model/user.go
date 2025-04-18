package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required"`
	Password string `json:"-" gorm:"unique;not null" validate:"required"`
}

func (u *User) EncryptPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedBytes)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
