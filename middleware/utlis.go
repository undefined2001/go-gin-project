package middleware

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"quizapp/models"
)

func ParseAndCreateQuiz(db *gorm.DB, filePath string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	quizTitle := ""
	questions := []models.Question{}

	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 7 {
			continue
		}

		title := row[0]
		if quizTitle == "" {
			quizTitle = title
		}

		q := models.Question{
			Question:    row[1],
			OptionOne:   row[2],
			OptionTwo:   row[3],
			OptionThree: row[4],
			OptionFour:  row[5],
		}
		fmt.Sscanf(row[6], "%d", &q.Answer)

		questions = append(questions, q)
	}

	quiz := models.Quiz{
		Title:     quizTitle,
		Questions: questions,
	}

	return db.Create(&quiz).Error
}
