// Package classification Kube Features API.
//
// the purpose of this application is to provide an application
// to use on Kubernetes and its features.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Raphael Amoedo<ralph.avalon@example.com> https://ralphavalon.github.io
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	api "kube_features/api/handlers"
	_ "kube_features/swagger-ui"
	"os"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// swagger:route GET /health HealthCheck health_check_without_db
	//
	// Get health check without checking database.
	//
	// Does a health check without checking database.
	//
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: healthCheckResponse
	//       503: healthCheckResponse
	router.Handle("/health", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(api.HealthCheck))).Methods("GET")
	// swagger:route POST /product Product productRequest
	//
	// Create a new product.
	//
	// Receives a product and saves on database.
	//
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: productResponse
	router.Handle("/product", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(api.CreateProduct))).Methods("POST")

	sh := http.StripPrefix("/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/").Handler(sh)

	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})

	h := handlers.CORS(corsOrigins, corsMethods)(router)

	http.ListenAndServe(":8081", h)
}
