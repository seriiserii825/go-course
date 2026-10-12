package utils

import "fmt"

func PrintShopTitle() {
	fmt.Println("=== SHOPPING CART ===")
}

func PrintPrices(prices []float64) {
	for i, price := range prices {
		fmt.Printf("%d. $%.2f\n", i+1, price)
	}
}

func PrintProductsLen(products []float64) {
	fmt.Printf("Total products: %d\n", len(products))
}

func PrintTotalPrice(prices []float64) {
	total := 0.0
	for _, price := range prices {
		total += price
	}
	fmt.Printf("Total price: $%.2f\n", total)
}

func PrintMostExpensive(prices []float64) {
	if len(prices) == 0 {
		fmt.Println("No products available.")
		return
	}
	mostExpensive := prices[0]
	for _, price := range prices {
		if price > mostExpensive {
			mostExpensive = price
		}
	}
	fmt.Printf("Most expensive: $%.2f\n", mostExpensive)
}
