package menu

import (
	"github.com/seriiserii825/go-course/homework/slice-map-struct/inventory"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/utils"
)

func ShowAllProducts(products []models.Product) {
	utils.PrintHeader("All products:")
	inventory.PrintProducts(products)
}
