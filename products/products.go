package products

import (
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
func (products Products) Create(product *Product) (*Product, *serverless.ServiceError) {

	if product == nil {
		return nil, serverless.Error("Invalid product", 401)
	}

	if product.Name == "" {
		return nil, serverless.Error("Name can't be empty", 401)
	}

	if product.Description == "" {
		return nil, serverless.Error("Description can't be empty", 401)
	}

	product.ID = uuid.Must(uuid.NewV4(), nil).String()
	product.CreatedAt = time.Now().String()

	if err := db.Add(product, os.Getenv("PRODUCTS_TABLE_NAME")); err != nil {
		return nil, serverless.Error(err.Error(), 501)
	}

	return product, nil
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
