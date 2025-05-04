package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quizapp/database"
	"quizapp/models"
)

func QuizIndex(c *gin.Context) {

	var quizzes []models.Quiz
	result := database.DB.Find(&quizzes)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.HTML(http.StatusOK, "quiz.html", gin.H{
		"Title":   "Quizzy",
		"Quizzes": quizzes,
	})
}

func GetParticipateQuiz(c *gin.Context) {
	quizID := c.Param("id")
	var questions []models.Question
	result := database.DB.Where("quiz_id = ?", quizID).Find(&questions)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var quiz models.Quiz
	result = database.DB.First(&quiz, quizID)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.HTML(http.StatusOK, "participate_quiz.html", gin.H{
		"Title":     "Participate Quiz",
		"Quiz":      quiz,
		"Questions": questions,
	})
}
