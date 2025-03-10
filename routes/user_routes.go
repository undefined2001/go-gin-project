package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	routes := router.Group("/user")
	routes.GET("/", controllers.GetUser)
	routes.POST("/login", controllers.UserLogin)
	routes.POST("/register", controllers.UserRegisterPOST)
	routes.GET("/register", controllers.UserRegisterGET)
}
