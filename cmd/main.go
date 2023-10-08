package main

import (
	"github.com/jswizzy/hello-api/handlers"
	"github.com/jswizzy/hello-api/handlers/rest"
	"github.com/jswizzy/hello-api/translation"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()
	service := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(service)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
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
