package controllers

import (
	"net/http"
	"quizapp/database"
	"quizapp/models"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "This is GetProduct Controller",
	})
}

func GetUserLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"Title": "Login", "csrf": csrf.GetToken(c)})

}

func PostUserLogin(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Check password
	if !user.CheckPassword(request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Return success response (you can add JWT token handling here)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func UserRegisterPOST(c *gin.Context) {

	firstName := c.PostForm("firstname")
	lastNmae := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := models.User{
		Username: firstName + " " + lastNmae,
		Email:    email,
		Password: password,
	}

	result := database.DB.Create(&user)

	print(user.Username)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func UserRegisterGET(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"csrf": csrf.GetToken(c),
	})
}
