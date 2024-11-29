package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go_di_example/pkg/logger"
	"net/http"
)

type ProductHandler struct {
	logger logger.Logger
}

func NewProductHandler(l logger.Logger) *ProductHandler {
	return &ProductHandler{
		logger: l,
	}
}

func (h *ProductHandler) HandlerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.list)
	r.Get("/{id}", h.get)
	r.Post("/", h.create)

	return r
}

func (h *ProductHandler) list(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Listing products")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Listing products")
	return
}

func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "get product")
	return
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Creating product")

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "create product")
	return
}
