package routes

import (
	"github.com/arjun-saseendran/event-project-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
}
