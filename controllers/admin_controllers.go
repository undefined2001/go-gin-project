package controllers

import (
	"net/http"
	"quizapp/database"
	"quizapp/models"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Admin | Dashboard",
	})
}

func AdminUserList(c *gin.Context) {
	var users []models.User
	database.DB.Where("role != ?", "admin").Find(&users)
	println("users", users)
	c.HTML(http.StatusOK, "user_list.html", gin.H{
		"Users": users,
		"Title": "Admin | User List",
	})
}

func GetViewQuiz(c *gin.Context) {
	var quizzes []models.Quiz

	result := database.DB.Find(&quizzes)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for _, quiz := range quizzes {
		println("quiz", quiz.Title)
	}

	c.HTML(http.StatusOK, "view_quiz_list.html", gin.H{
		"Title":   "Admin | Quiz List",
		"Quizzes": quizzes,
	})
}


func GetAddQuize(c *gin.Context) {
	c.HTML(http.StatusOK, "add_quiz.html", gin.H{
		"Title": "Add Quiz",
		"csrf":  csrf.GetToken(c),
	})

}

func PostAddQuiz(c *gin.Context) {
	// Get the form values
	quizTitle := c.PostForm("title")

	quiz := &models.Quiz{
		Title: quizTitle,
	}
	println("quiz", quiz.Title)
	database.DB.Create(quiz)

	c.Redirect(http.StatusSeeOther, "/admin/add_quiz")
}
