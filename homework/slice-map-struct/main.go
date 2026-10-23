package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/inventory"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/utils"
)

func main() {
	products := []models.Product{
		{ID: 1, Name: "Product 1", Price: 10.0, Quantity: 5, Category: "Category A"},
		{ID: 2, Name: "Product 2", Price: 20.0, Quantity: 3, Category: "Category B"},
		{ID: 3, Name: "Product 3", Price: 15.0, Quantity: 8, Category: "Category A"},
		{ID: 4, Name: "Product 4", Price: 30.0, Quantity: 2, Category: "Category C"},
		{ID: 5, Name: "Product 5", Price: 25.0, Quantity: 6, Category: "Category B"},
	}

	inventory.PrintProducts(products)
	searchedProductId := 3
	product2 := inventory.SearchProductById(products, searchedProductId)
	if product2 != nil {
		utils.PrintHeader(fmt.Sprintf("Product with ID %d:", product2.ID))
		pp.Println(product2)
	} else {
		utils.PrintErrorMessage(fmt.Sprintf("Product with ID %d not found", searchedProductId))
	}

	minPrice := 100.0
	filteredProducts := inventory.FilterByPriceMin(products, minPrice)
	if len(filteredProducts) > 0 {
		utils.PrintHeader(fmt.Sprintf("Products with price >= %.2f, Count: %d", minPrice, len(filteredProducts)))
		inventory.PrintProducts(filteredProducts)
	} else {
		utils.PrintErrorMessage(fmt.Sprintf("No products found with price >= %.2f", minPrice))
	}

	moreExpensiveProduct, ok := inventory.GetMoreExpensiveProduct(products)
	if ok {
		utils.PrintHeader("More Expensive product:")
		pp.Println(moreExpensiveProduct)
	} else {
		utils.PrintErrorMessage("No products found in the inventory")
	}

	lessExpensiveProduct, ok := inventory.GetLessExpensiveProduct(products)
	if ok {
		utils.PrintHeader("Less Expensive product:")
		pp.Println(lessExpensiveProduct)
	} else {
		utils.PrintErrorMessage("No products found in the inventory")
	}

	storeTotalPrice := inventory.GetStoreTotalPrice(products)
	utils.PrintHeader("Store total price: ")
	fmt.Printf("%.2f$\n", storeTotalPrice)
}
