package routes

import (
	"crud/pkg/controller"

	"github.com/gorilla/mux"
)

func EmployeeRoutes(r *mux.Router) {
	r.HandleFunc("/users", controller.Users).Methods("GET")
	r.HandleFunc("/user/{id}", controller.User).Methods("GET")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}",controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

}
