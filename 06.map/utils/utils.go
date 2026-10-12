package utils

import "fmt"

func PrintTitle() {
	println("=== PRODUCT INVENTORY ===")
}

func printAllProducts(products map[string]int) {
	for product, quantity := range products {
		result := fmt.Sprintf("%s: %d", product, quantity)
		fmt.Println(result)
	}
}

func AddProduct(name string, count int, products map[string]int) map[string]int {
	products[name] = count
	return products
}

func ChangeQuantity(name string, count int, products map[string]int) map[string]int {
	if _, exists := products[name]; exists {
		currentCount := products[name]
		products[name] = currentCount + count
	}
	return products
}

func DeleteProduct(name string, products map[string]int) map[string]int {
	delete(products, name)
	return products
}

func isProductExist(name string, products map[string]int) bool {
	_, exists := products[name]
	return exists
}

func getMostStockedProduct(products map[string]int) (string, int) {
	var mostStockedProduct string
	var maxQuantity int

	for product, quantity := range products {
		if quantity > maxQuantity {
			maxQuantity = quantity
			mostStockedProduct = product
		}
	}

	return mostStockedProduct, maxQuantity
}

func getAllProductsQuantity(products map[string]int) int {
	totalQuantity := 0
	for _, quantity := range products {
		totalQuantity += quantity
	}
	return totalQuantity
}

func FormatResult(products map[string]int) {
	fmt.Println(` === PRODUCT INVENTORY ===`)
	fmt.Println(`Products:`)
	printAllProducts(products)
	mostStockedProduct, quantity := getMostStockedProduct(products)

	fmt.Printf("Monitor exists: %t\n", isProductExist("Monitor", products))
	fmt.Printf("Most stocked product: %s (%d)\n", mostStockedProduct, quantity)
	fmt.Printf("Total stock: %d\n", getAllProductsQuantity(products))
	fmt.Printf("Total products count: %d\n", len(products))
}
