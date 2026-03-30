package main

import (
	"github.com/arjun-saseendran/event-project-go-gin/db"
	"github.com/arjun-saseendran/event-project-go-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect database
	db.InitDB()

	// config sever
	server := gin.Default()

	// config routes
	routes.RegisterEventRoutes(server)
	routes.RegisterUserRoutes(server)

	// config port
	server.Run(":3000")

}
