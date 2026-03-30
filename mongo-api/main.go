package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/arjun-saseendran/mongoapi/routes"
)

func main() {
	r := router.Router()
	fmt.Println("Server is getting started..")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Server running on port 3000")
}
