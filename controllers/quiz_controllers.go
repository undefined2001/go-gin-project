package controllers

import (
	"net/http"
	"quizapp/database"
	"quizapp/models"
	"github.com/gin-gonic/gin"
)

func QuizIndex(c *gin.Context) {
	quizID := c.Param("id")
	var quiz models.Quiz

	// Fetch quiz with associated questions
	if err := database.DB.Preload("Questions").First(&quiz, quizID).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.HTML(http.StatusOK, "quiz.html", gin.H{
		"Title": "Quiz | " + quiz.Title,
		"Quiz":  quiz,
	})
}
