package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Question    string `json:"question" binding:"required"`
	OptionOne   string `json:"option_one" binding:"required"`
	OptionTwo   string `json:"option_two" binding:"required"`
	OptionThree string `json:"option_three" binding:"required"`
	OptionFour  string `json:"option_four" binding:"required"`
	Answer      int    `json:"answer" binding:"required"`
}

func (q *Quiz) CreateQuiz(db *gorm.DB) error {
	err := db.Create(q).Error
	if err != nil {
		return err
	}
	return nil
}
func (q *Quiz) GetAllQuizzes(db *gorm.DB) ([]Quiz, error) {
	var quizzes []Quiz
	err := db.Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}