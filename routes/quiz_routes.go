package routes

import (
	"github.com/gin-gonic/gin"
	"quizapp/controllers"
)

// Register task-related routes
func QuizRoutes(router *gin.Engine) {
	task_group := router.Group("/")
	task_group.GET("/", controllers.QuizIndex)
	task_group.GET("/:id", controllers.GetParticipateQuiz)
}
