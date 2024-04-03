package main

import (
	"Go-JWT/configs"
	"Go-JWT/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
