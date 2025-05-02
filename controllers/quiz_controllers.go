package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QuizIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", nil)
}
