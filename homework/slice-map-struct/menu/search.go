package menu

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/inventory"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/utils"
)

func ShowProductById(products []models.Product, id int) {
	product := inventory.SearchProductById(products, id)
	if product == nil {
		utils.PrintErrorMessage(fmt.Sprintf("Product with ID %d not found", id))
		return
	}
	utils.PrintHeader(fmt.Sprintf("Product with ID %d:", product.ID))
	pp.Println(product)
}

func ShowProductsByPriceMin(products []models.Product, minPrice float64) {
	filteredProducts := inventory.FilterByPriceMin(products, minPrice)
	if len(filteredProducts) == 0 {
		utils.PrintErrorMessage(fmt.Sprintf("No products found with price >= %.2f", minPrice))
		return
	}
	utils.PrintHeader(fmt.Sprintf("Products with price >= %.2f, Count: %d", minPrice, len(filteredProducts)))
	inventory.PrintProducts(filteredProducts)
}
