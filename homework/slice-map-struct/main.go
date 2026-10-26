package main

import (
	"github.com/seriiserii825/go-course/homework/slice-map-struct/menu"
	"github.com/seriiserii825/go-course/homework/slice-map-struct/models"
)

func main() {
	products := []models.Product{
		{ID: 1, Name: "Product 1", Price: 10.0, Quantity: 5, Category: "Category A"},
		{ID: 2, Name: "Product 2", Price: 20.0, Quantity: 3, Category: "Category B"},
		{ID: 3, Name: "Product 3", Price: 15.0, Quantity: 8, Category: "Category A"},
		{ID: 4, Name: "Product 4", Price: 30.0, Quantity: 2, Category: "Category C"},
		{ID: 5, Name: "Product 5", Price: 25.0, Quantity: 6, Category: "Category B"},
	}

	menu.ShowAllProducts(products)
	menu.ShowProductById(products, 3)
	menu.ShowProductsByPriceMin(products, 100.0)
	menu.ShowMoreExpensiveProduct(products)
	menu.ShowLessExpensiveProduct(products)
	menu.ShowStoreTotalPrice(products)
}
