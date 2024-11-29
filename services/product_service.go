package services

import (
	"go_di_example/models"
	"go_di_example/pkg/logger"
	"go_di_example/stores"
)

type ProductService interface {
	List() ([]models.Product, error)
	Get(id int) (models.Product, error)
	Create(models.Product) error
}

type productService struct {
	logger logger.Logger
	store  stores.ProductStore
}

func NewProductService(l logger.Logger, s stores.ProductStore) ProductService {
	return &productService{logger: l, store: s}
}

func (p *productService) List() ([]models.Product, error) {
	return p.store.List()
}

func (p *productService) Get(id int) (models.Product, error) {
	return p.store.Get(id)
}

func (p *productService) Create(product models.Product) error {
	return p.store.Create(product)
}
