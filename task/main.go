package main

import (
	"net/http"
	"second-project/config"
	"second-project/routes"
)

func main() {
	router := routes.MountRoutes()
	config.InitDB()

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: router,
	}

	server.ListenAndServe()
}
