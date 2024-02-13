package router

import (
	"tryUrl/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/createshorturl", controller.ShortenURL).Methods("POST")
	router.HandleFunc("/{shortURL}", controller.RedirectURL).Methods("GET")
	return router
}