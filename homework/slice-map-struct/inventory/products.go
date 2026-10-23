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
