package stores

import (
	"database/sql"
	"errors"
	"go_di_example/models"
	"go_di_example/pkg/logger"
	"time"
)

type ProductStore interface {
	List() ([]models.Product, error)
	Get(id int) (models.Product, error)
	Create(product models.Product) error
}

type productStore struct {
	logger logger.Logger
	db     *sql.DB
}

func NewProductStore(l logger.Logger, db *sql.DB) ProductStore {
	return &productStore{logger: l, db: db}
}

func (p *productStore) List() ([]models.Product, error) {
	rows, err := p.db.Query("SELECT id, name, price, created_at FROM products")
	if err != nil {
		p.logger.Error("Error querying products: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	var products []models.Product

	for rows.Next() {
		var product models.Product

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt); err != nil {
			p.logger.Error("Error scanning product: " + err.Error())
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		p.logger.Error("Error scanning products: " + err.Error())
		return nil, err
	}
	return products, nil
}

func (p *productStore) Get(id int) (models.Product, error) {
	var product models.Product
	row := p.db.QueryRow("select id, name, price, created_at from products where id = ?", id)
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.logger.Error("Product not found with id: " + string(rune(id)))
			return product, nil //there are no errors, just the product was not found
		}
		p.logger.Error("Error querying product: " + err.Error())
		return product, err
	}
	return product, nil
}

func (p *productStore) Create(product models.Product) error {
	product.CreatedAt = time.Now().Format(time.RFC3339)

	_, err := p.db.Exec("INSERT INTO products (name, price, created_at) VALUES (?, ?, ?)", product.Name, product.Price, product.CreatedAt)
	if err != nil {
		p.logger.Error("Error inserting product: " + err.Error())
		return err
	}

	return nil
}
