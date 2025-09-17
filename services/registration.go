package services

import (
	"net/http"

	"github.com/car-api/model"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user model.Users

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error while binding the json."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error while saving the data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
