package main

import (
	"net/http"
	"quizapp/database"
	"quizapp/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"quizapp/middleware"
)

func main() {

	port := ":8080"
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	store := cookie.NewStore([]byte("super-secret-key"))

	r.Use(sessions.Sessions("Mysessions", store))

	r.Use(middleware.AuthMiddleware()) // * Auth Middleware

	r.Use(csrf.Middleware(
		csrf.Options{
			Secret: "MySecretKey",
			ErrorFunc: func(c *gin.Context) {
				c.String(http.StatusForbidden, "CSRF Token Mismatched")
				c.Abort()
			},
		}))

	database.ConnectDB() //* Connecting to the Database

	r.Static("/static", "./static")

	routes.UserRoutes(r)  // Registering User Routes
	routes.QuizRoutes(r)  // Registering Quiz Routes
	routes.AdminRoutes(r) //Registering Admin Routes

	r.Run(port)

}
