package controllers

import (
	"net/http"
"github.com/arjun-saseendran/event-project-go-gin/models"
	"github.com/arjun-saseendran/event-project-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not signup user!"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})

}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login!"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorised user!"})
		return
	}
	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successfull!", "token": token})

}
