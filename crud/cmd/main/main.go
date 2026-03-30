package main

import (
	"fmt"
	"log"
	"net/http"

	"crud/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.EmployeeRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server running on port: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
