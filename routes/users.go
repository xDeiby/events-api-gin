package routes

import (
	"fmt"
	"net/http"

	"example.com/gin-project/models"
	"example.com/gin-project/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		
		return
	}

	_, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		
		return	
	}

	var user models.User

	err = context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save data",
		})
	}

	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed",
		})
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Authentication failed to generate token",
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}