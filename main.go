package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"time"

	api "kube_features/api/handlers"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Server *http.Server

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", api.HealthCheck).Methods("GET")

	// Configure CORS (allow all origins, and some methods for now)
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})

	h := handlers.CORS(corsOrigins, corsMethods)(router)
	s := &http.Server{Addr: "0.0.0.0:" + strconv.Itoa(8081), Handler: h, ReadTimeout: 7 * time.Second, WriteTimeout: 7 * time.Second}
	Server = s

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := s.ListenAndServe(); err != nil {
			//log.Fatalf("• Error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	Server.Shutdown(ctx)
	os.Exit(0)
}
