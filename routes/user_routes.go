package routes

import (
	"quizapp/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	routes := router.Group("/user")
	routes.GET("/", controllers.GetUser)
	routes.GET("/login", controllers.GetUserLogin)
	routes.POST("/login", controllers.PostUserLogin)
	routes.POST("/register", controllers.UserRegisterPOST)
	routes.GET("/register", controllers.UserRegisterGET)
}
