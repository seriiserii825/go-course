package inventory

import "github.com/seriiserii825/go-course/homework/slice-map-struct/models"

func GroupByCategory(products []models.Product) map[string][]models.Product {
	result := make(map[string][]models.Product)

	for _, product := range products {
		result[product.Category] = append(result[product.Category], product)
	}
	return result
}
