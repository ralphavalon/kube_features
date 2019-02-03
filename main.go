package main

import (
	api "kube_features/api/handlers"
	"os"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/health", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(api.HealthCheck))).Methods("GET")
	router.Handle("/product", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(api.CreateProduct))).Methods("POST")

	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})

	h := handlers.CORS(corsOrigins, corsMethods)(router)

	http.ListenAndServe(":8081", h)
}
