package routes

import (
	"quizapp/controllers"
	"github.com/gin-gonic/gin"
)

// Register task-related routes
func QuizRoutes(router *gin.Engine) {
	task_group := router.Group("/quiz")
	task_group.GET("/", controllers.QuizIndex)
}
