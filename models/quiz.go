package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Title     string     `json:"title"`
	Questions []Question `gorm:"foreignKey:QuizID;constraint:OnDelete:CASCADE;" json:"questions"`
}

type Question struct {
	gorm.Model
	QuizID      uint
	Question    string `json:"question" binding:"required"`
	OptionOne   string `json:"option_one" binding:"required"`
	OptionTwo   string `json:"option_two" binding:"required"`
	OptionThree string `json:"option_three" binding:"required"`
	OptionFour  string `json:"option_four" binding:"required"`
	Answer      int    `json:"answer" binding:"required"`
}

func (q *Question) CreateQuestion(db *gorm.DB) error {
	err := db.Create(q).Error
	if err != nil {
		return err
	}
	return nil
}
func (q *Question) GetAllQuestion(db *gorm.DB) ([]Quiz, error) {
	var questions []Quiz
	err := db.Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
