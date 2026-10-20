package products

func CreateProduct(id int, name string, price float64, quantity int, category string) Product {
	return Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}
}
