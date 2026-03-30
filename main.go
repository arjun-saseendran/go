package main

import (
	"log"
	"workout/db"
	"workout/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	// config db
	db.InitDB()
	
	// config server
	server := gin.Default()

	// config routes
	routes.UserRoutes(server)

	// config port & run
	err := server.Run(":3000")
	if err != nil {
		log.Fatal("Failed to run server!")
	}

}
