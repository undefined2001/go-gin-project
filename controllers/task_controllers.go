package controllers

import (
	"net/http"
	"todolist/database"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

// Mock task li

// Get all tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	// Perform the query and populate the 'tasks' slice
	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve tasks",
		})
		return
	}

	// Return tasks as HTML response
	c.HTML(http.StatusOK, "index.html", gin.H{
		"tasks": tasks,
	})
}

// Create a new task
func CreateTask(c *gin.Context) {
	task := models.Task{Title: "This is a Test Task", Status: "Pending"}

	result := database.DB.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
	}

	c.JSON(http.StatusCreated, task)
}
