package routes

import (
	"workout/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {
	server.GET("/users", controllers.GetUsers)
	server.GET("/user/:id", controllers.GetUser)
	server.POST("/user", controllers.CreateUser)
	server.PUT("/user/:id", controllers.UpdateUser)
	server.DELETE("/user/:id", controllers.DeleteUser)
}
