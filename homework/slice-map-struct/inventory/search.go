package inventory

import "github.com/seriiserii825/go-course/homework/slice-map-struct/models"

func FilterByPriceMin(products []models.Product, minPrice float64) []models.Product {
	var filteredProducts []models.Product
	for _, product := range products {
		if product.Price >= minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
