package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"workout/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	// config user
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Check required fields!", "details": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Falild to save user to database!", "error": err.Error()})
		return
	}
	// response to client
	context.JSON(http.StatusCreated, gin.H{"message": "User created successflly!", "user": user})
}

func GetUser(context *gin.Context) {
	// get id from url
	idString := context.Param("id")

	// parse id to int
	id, err := strconv.Atoi(idString)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user id format! It must be a number!"})
		return
	}

	// find the user
	user, err := models.GetUserByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response user data
	context.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(context *gin.Context) {
	// call func for users data
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrive data from database!", "error": err.Error()})
		return
	}
	// reponse to client
	context.JSON(http.StatusOK, gin.H{"users": users})
}

func UpdateUser(context *gin.Context) {
	// get id from url
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	//config user
	var updatedUser models.User

	if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data!"})
		return
	}

	// set url id to user id
	updatedUser.ID = id

	if err := updatedUser.Update(); err != nil {
		if err.Error() == fmt.Sprintf("user with id %d not found", id) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not update user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user updated successfully!", "user": updatedUser})
}

func DeleteUser(context *gin.Context) {
	// get id from url
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id formart!"})
		return
	}
	if err := models.Delete(id); err != nil {
		if err.Error() == fmt.Sprintf("user with id %d not found!", id) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user!"})
		return
	}
	// success response to client
	context.JSON(http.StatusOK, gin.H{"message": "user deleted successfully!"})
}
