package routes

import (
	"github.com/arjun-saseendran/event-project-go-gin/controllers"
	"github.com/arjun-saseendran/event-project-go-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	server.GET("/events", controllers.GetEvents)
	server.GET("/event/:id", controllers.GetEvent)
	authenticated.POST("/event", controllers.CreateEvent)
	authenticated.PUT("/event/:id", controllers.UpdateEvent)
	authenticated.DELETE("/event/:id", controllers.DeleteEvent)
	authenticated.POST("/event/:id/register", controllers.RegisterForEvent)
	authenticated.DELETE("/event/:id/cancel", controllers.CancelForRegistration)

}
