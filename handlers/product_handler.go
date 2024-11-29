package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go_di_example/models"
	"go_di_example/pkg/logger"
	"go_di_example/services"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	logger         logger.Logger
	productService services.ProductService
}

func NewProductHandler(l logger.Logger, service services.ProductService) *ProductHandler {
	return &ProductHandler{
		logger:         l,
		productService: service,
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

	products, err := h.productService.List()
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": products})
	return
}

func (h *ProductHandler) get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	product, err := h.productService.Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": product})
	return
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		h.logger.Error("invalid input" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	err := h.productService.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "created")
	return
}
