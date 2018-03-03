package products

import (
	"errors"
	"os"
	"time"

	"github.com/petargk/serverless.products/lib"
	"github.com/satori/go.uuid"
)

// Products service
type Products struct{}

// Product struct
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// List products struct
type List struct {
	Products []Product `json:"products"`
}

var db serverless.Dynamodb

// Create product
func (products Products) Create(product *Product) (*Product, int, error) {

	if product == nil {
		return nil, 401, errors.New("Invalid product")
	}

	if product.Name == "" {
		return product, 401, errors.New("Name can't be empty")
	}

	if product.Description == "" {
		return product, 401, errors.New("Description can't be empty")
	}

	product.ID = uuid.Must(uuid.NewV4(), nil).String()
	product.CreatedAt = time.Now().String()

	if err := db.Add(product, os.Getenv("PRODUCTS_TABLE_NAME")); err != nil {
		return nil, 501, err
	}

	return product, 201, nil
}

// Update product
func (products Products) Update(product Product) (Product, error) {

	return product, nil
}

// Delete product
func (products Products) Delete(id string) (bool, error) {

	return true, nil
}

// GetAll products
//func (products Products) GetAll() (List, error) {
//
//	return nil, nil
//}
