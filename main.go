package main

import (
	"net/http"

	"github.com/Rojas-Andres/project-api-rest-go/db"
	"github.com/Rojas-Andres/project-api-rest-go/routes"
	"github.com/gorilla/mux"
)
func main() {

	db.DBConnection()
	router := mux.NewRouter()
	// router.HandleFunc("/api/v1/health", HealthCheck).Methods("GET")
	
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}