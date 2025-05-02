package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		session := sessions.Default(c)
		user := session.Get("user")
		print(user)
		// if user == nil {
		// 	c.Redirect(http.StatusFound, "user/login")
		// 	c.Abort()
		// 	return
		// }
		// if user.IsAdmin() {
		// 	c.Redirect(http.StatusFound, "/admin")
		// 	c.Abort()
		// 	return
		// } else {
		// 	//Here Goes else Part
		// }
		c.Next()
	}
}
