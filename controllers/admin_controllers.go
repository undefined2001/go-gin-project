package controllers

import (
	"net/http"
	"quizapp/database"
	"quizapp/models"
	"strconv"

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

func GetAddQuize(c *gin.Context) {
	c.HTML(http.StatusOK, "add_quiz.html", gin.H{
		"Title": "Add Quiz",
		"csrf":  csrf.GetToken(c),
	})

}

func PostAddQuiz(c *gin.Context) {
	// Get the form values
	quizQuestion := c.PostForm("question")
	quizOptionOne := c.PostForm("option1")
	quizOptionTwo := c.PostForm("option2")
	quizOptionThree := c.PostForm("option3")
	quizOptionFour := c.PostForm("option4")
	quizAnswer, _ := strconv.Atoi(c.PostForm("answer"))

	quiz := &models.Quiz{
		Question:    quizQuestion,
		OptionOne:   quizOptionOne,
		OptionTwo:   quizOptionTwo,
		OptionThree: quizOptionThree,
		OptionFour:  quizOptionFour,
		Answer:      quizAnswer,
	}

	if err := quiz.CreateQuiz(database.DB); err != nil {
		println("Quiz Creation Failed.")
	}

	c.Redirect(http.StatusSeeOther, "/admin/add_quiz")
}
