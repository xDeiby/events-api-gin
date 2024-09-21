package routes

import (
	"fmt"
	"net/http"

	"example.com/gin-project/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

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