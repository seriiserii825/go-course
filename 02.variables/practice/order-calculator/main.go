package main

import "fmt"

// Напиши небольшой расчёт заказа интернет-магазина, аналогичного твоему Angular Shop.
// productName := "Gaming Keyboard"
// price := 89.99
// quantity := 3
// const discount = 15.0
// const vat = 20.0
// Нужно самостоятельно вычислить стоимость заказа, скидку 15%, стоимость после скидки, НДС 20% от суммы после скидки и окончательную стоимость с НДС.
// Ожидаемый вывод:
// === Order Summary ===
// Product: Gaming Keyboard
// Price: $89.99
// Quantity: 3
// Subtotal: $269.97
// Discount (15%): $40.50
// After discount: $229.47
// VAT (20%): $45.89
// Total: $275.36
// Для расчётов используй float64. Позже изучим, почему для настоящих денежных операций лучше использовать целые значения минимальных денежных единиц или decimal-библиотеки.
func main() {
	productName := "Gaming Keyboard"
	price := 89.99
	quantity := 3
	const discount = 15.0
	const vat = 20.0
	const delimiter = "=============="
	subtotal := float64(quantity) * price
	discountPercent := (subtotal * discount) / 100
	afterDiscount := subtotal - discountPercent
	vatPercent := (afterDiscount * vat) / 100
	total := afterDiscount + vatPercent

	fmt.Println("Product Name: ", productName)
	fmt.Printf("Price: $%.2f\n", price)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Discount ($%.f): %.2f\n", discount, discountPercent)
	fmt.Printf("After discount: $%.2f\n", afterDiscount)
	fmt.Printf("VAT (%.f%%): %.2f\n", vat, vatPercent)
	fmt.Printf("Total: %.2f\n", total)
}
