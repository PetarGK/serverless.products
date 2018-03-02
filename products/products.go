package products

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
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

// Create product
func (products Products) Create(p Product) (Product, error) {
	var (
		id        = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("PRODUCTS_TABLE_NAME"))
	)

	return p, nil
}

// Update product
func (products Products) Update(p Product) (Product, error) {

	return p, nil
}

// Delete product
func (products Products) Delete(id string) (bool, error) {

	return true, nil
}

// GetAll products
func (products Products) GetAll() (List, error) {

	return nil, nil
}
