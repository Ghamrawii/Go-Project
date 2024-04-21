package routes

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/events/models"
	"example.com/events/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	// Bind JSON to user struct
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request"})
		return
	}

	// Attempt to save user
	if err := user.Save(); err != nil {
		fmt.Println("Error saving user:", err)
		// Check for unique constraint error on email field
		if strings.Contains(err.Error(), "UNIQUE constraint failed") && strings.Contains(err.Error(), "users.email") {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Email already in use. Please use a different email."})
		} else {
			// Log unexpected errors and respond accordingly
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		}
		return
	}
	// Successfully created user
	context.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

func login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request"})
		return
	}

	err = user.ValidLogin()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user."})
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfully", "token": token})
}
