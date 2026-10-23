package inventory

import "github.com/seriiserii825/go-course/homework/slice-map-struct/models"

func GetMoreExpensiveProduct(products []models.Product) models.Product {
	expensiveProduct := products[0]

	for _, product := range products {
		if product.Price > expensiveProduct.Price {
			expensiveProduct = product
		}
	}
	return expensiveProduct
}

func GetLessExpensiveProduct(products []models.Product) models.Product {
	expensiveProduct := products[0]

	for _, product := range products {
		if product.Price < expensiveProduct.Price {
			expensiveProduct = product
		}
	}
	return expensiveProduct
}
