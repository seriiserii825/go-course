package main

import "github.com/seriiserii825/go-course/06.map/utils"

func main() {
	products := map[string]int{
		"Laptop":   12,
		"Mouse":    35,
		"Keyboard": 20,
		"Monitor":  8,
		"Headset":  15,
	}
	utils.PrintTitle()
	products = utils.AddProduct("Webcam", 5, products)
	products = utils.AddProduct("Mouse", 15, products)
	products = utils.ChangeQuantity("Laptop", -3, products)
	products = utils.DeleteProduct("Headset", products)
	utils.FormatResult(products)
}
