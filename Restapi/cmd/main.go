package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/Employees", handlers.GetAllEmployees).Methods(http.MethodGet)
	router.HandleFunc("/Employees/{id}", handlers.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/Employees", handlers.AddEmployee).Methods(http.MethodPost)
	router.HandleFunc("/Employees/{id}", handlers.UpdateEmployee).Methods(http.MethodPut)
	router.HandleFunc("/Employees/{id}", handlers.DeleteEmployee).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4001", router)
}
