package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"-" binding:"required"`
	Role     string `json:"role" gorm:"default:'user'"`
	Banned   bool   `json:"banned" gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashBytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
func (u *User) IsBanned() bool {
	return u.Banned
}
