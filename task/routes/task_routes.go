package routes

import (
	"second-project/controllers"

	"github.com/gin-gonic/gin"
)

func MountRoutes() *gin.Engine {
	router := gin.Default()

	taskRoutes := router.Group("/task")
	{
		taskRoutes.POST("/", controllers.SaveTask)
		taskRoutes.GET("/", controllers.GetTasks)

	}

	router.NoRoute(controllers.NotFound)

	return router
}
