package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

// Register task-related routes
func TaskRoutes(router *gin.Engine) {
	task_group := router.Group("/tasks")
	task_group.GET("/", controllers.GetTasks)
	task_group.GET("/create", controllers.CreateTask)
}
