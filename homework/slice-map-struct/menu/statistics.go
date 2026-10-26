package menu

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/inventory"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/utils"
)

func ShowMoreExpensiveProduct(products []models.Product) {
	product, ok := inventory.GetMoreExpensiveProduct(products)
	if !ok {
		utils.PrintErrorMessage("No products found in the inventory")
		return
	}
	utils.PrintHeader("More Expensive product:")
	pp.Println(product)
}

func ShowLessExpensiveProduct(products []models.Product) {
	product, ok := inventory.GetLessExpensiveProduct(products)
	if !ok {
		utils.PrintErrorMessage("No products found in the inventory")
		return
	}
	utils.PrintHeader("Less Expensive product:")
	pp.Println(product)
}

func ShowStoreTotalPrice(products []models.Product) {
	utils.PrintHeader("Store total price: ")
	fmt.Printf("%.2f$\n", inventory.GetStoreTotalPrice(products))
}
