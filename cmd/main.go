package main

import (
	"github.com/go-chi/chi/v5"
	"go_di_example/db"
	"go_di_example/handlers"
	"go_di_example/pkg/logger"
	"go_di_example/services"
	"go_di_example/stores"
	"log"
	"net/http"
)

func main() {

	dbConn := db.Connect("go_di_example.sqlite")
	db.CreateTable(dbConn)

	r := chi.NewRouter()
	var logrusLog logger.Logger = &logger.LogrusLogger{}

	productStore := stores.NewProductStore(logrusLog, dbConn)
	productService := services.NewProductService(logrusLog, productStore)
	productHandler := handlers.NewProductHandler(logrusLog, productService)

	r.Mount("/products", productHandler.HandlerRoutes())

	log.Println("======== Starting server on :8080 =======")
	log.Fatal(http.ListenAndServe(":8080", r))
}
