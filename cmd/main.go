package main

import (
	"github.com/jswizzy/hello-api/handlers"
	"github.com/jswizzy/hello-api/handlers/rest"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	log.Printf("listening on %s\n", addr)

	log.Fatal(server.ListenAndServe())
}
