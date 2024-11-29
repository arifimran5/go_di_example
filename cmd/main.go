package main

import (
	"github.com/go-chi/chi/v5"
	"go_di_example/handlers"
	"go_di_example/pkg/logger"
	"log"
	"net/http"
)

func init() {
	// connect to database here
}

func main() {
	r := chi.NewRouter()

	var l logger.Logger = &logger.LogrusLogger{}
	productHandler := handlers.NewProductHandler(l)

	r.Mount("/products", productHandler.HandlerRoutes())

	log.Println("======== Starting server on :8080 =======")
	log.Fatal(http.ListenAndServe(":8080", r))
}
