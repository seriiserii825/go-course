package models

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
	Category string
}

func NewProduct(id int, name string, price float64, quantity int, category string) Product {
	return Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}
}
