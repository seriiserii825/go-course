package main

import "strconv"
import "fmt"

func main() {
	// int → string
	age := strconv.Itoa(25)
	fmt.Println("Age:", age)

	// float64 → string
	price := strconv.FormatFloat(99.99, 'f', 2, 64)
	fmt.Println("Price:", price)
}
