package controllers

import (
	"net/http"
	"second-project/models"

	"github.com/gin-gonic/gin"
)

func SaveTask(ctx *gin.Context) {
	var payload models.PostTaskPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body!"})
		return
	}
	id, err := models.TaskRepositary.SaveTaskQuery(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "message": id})

}

func GetTasks(ctx *gin.Context) {

	tasks, err := models.TaskRepositary.ReadTaskQuery()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task successfully fetched!", "task": tasks})
}

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Url not found!"})

}
