package inventory

import "github.com/seriiserii825/go-course/homework/slice-map-struct/models"

func GetMoreExpensiveProduct(products []models.Product) (models.Product, bool) {
	if len(products) == 0 {
		return models.Product{}, false
	}

	expensiveProduct := products[0]

	for _, product := range products {
		if product.Price > expensiveProduct.Price {
			expensiveProduct = product
		}
	}
	return expensiveProduct, true
}

func GetLessExpensiveProduct(products []models.Product) (models.Product, bool) {
	if len(products) == 0 {
		return models.Product{}, false
	}
	cheapestProduct := products[0]

	for _, product := range products {
		if product.Price < cheapestProduct.Price {
			cheapestProduct = product
		}
	}
	return cheapestProduct, true
}

func GetStoreTotalPrice(products []models.Product) float64 {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += float64(product.Quantity) * product.Price
	}
	return totalPrice
}
