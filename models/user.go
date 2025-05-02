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

type Score struct {
	gorm.Model

	UserID uint `gorm:"uniqueIndex:idx_user_quiz" json:"user_id"`
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	QuizID uint `gorm:"uniqueIndex:idx_user_quiz" json:"quiz_id"`
	Quiz   Quiz `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"quiz"`

	Points int `json:"score"`
}

func (s *Score) CreateScore(db *gorm.DB) error {
	err := db.Create(s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Score) GetUser(db *gorm.DB) (*User, error) {
	var user User
	err := db.First(&user, s.UserID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Score) GetQuiz(db *gorm.DB) (*Quiz, error) {
	var quiz Quiz
	err := db.First(&quiz, s.QuizID).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}