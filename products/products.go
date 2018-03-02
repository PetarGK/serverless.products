package products

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
func (products Products) Create(p Product) {

}

// Update product
func (products Products) Update(p Product) {

}

// Delete product
func (products Products) Delete(id string) {

}

// GetAll products
func (products Products) GetAll() {

}
