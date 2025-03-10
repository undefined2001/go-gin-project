package main

import (
	"net/http"
	"todolist/database"
	"todolist/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func main() {

	port := ":8080"
	r := gin.Default()

	store := cookie.NewStore([]byte("super-secret-key"))

	r.Use(sessions.Sessions("Mysessions", store))

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
	r.LoadHTMLGlob("./templates/*") //* Setting Folder for Templates

	routes.TaskRoutes(r) // Registering Task Routes
	routes.UserRoutes(r) // Registering User Routes

	r.Run(port)

}
