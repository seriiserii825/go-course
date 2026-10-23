package inventory

import (
	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
)

func PrintProducts(products []models.Product) {
	for _, product := range products {
		pp.Println(product)
	}
}

func SearchProductById(products []models.Product, id int) *models.Product {
	for i, _ := range products {
		if products[i].ID == id {
			return &products[i]
		}
	}
	return nil
}
