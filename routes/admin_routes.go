package routes

import (
	"github.com/gin-gonic/gin"
	"quizapp/controllers"
)

func AdminRoutes(router *gin.Engine) {
	admin_routes := router.Group("/admin")
	admin_routes.GET("/", controllers.AdminIndex)
	admin_routes.GET("/user_list", controllers.AdminUserList)

	//Quiz Related Routes
	admin_routes.GET("/view_quiz", controllers.GetViewQuiz)
	admin_routes.GET("/add_quiz", controllers.GetAddQuize)
	admin_routes.POST("/add_quiz", controllers.PostAddQuiz)
	admin_routes.POST("/upload", controllers.PostUploadQuizExcelSheet)
	admin_routes.GET("/upload", controllers.GetUploadQuizExcelSheet)
}
