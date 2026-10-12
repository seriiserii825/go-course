package main

import (
	"fmt"
	"strconv"
)

// 02. Shopping Cart Calculator
//
// Ты создаёшь программу расчёта стоимости заказа
// для интернет-магазина.
//
// Исходные данные:
//
// Задание:
// 1. Преобразуй price в float64.
// 2. Преобразуй quantity в int.
// 3. Преобразуй discount в float64.
// 4. Вычисли общую стоимость товаров.
// 5. Рассчитай сумму скидки 15%.
// 6. Рассчитай итоговую стоимость.
// 7. Преобразуй итоговую стоимость в строку
//    через strconv.FormatFloat().
// 8. Выведи результат через fmt.Printf().
//
// Ожидаемый результат:
// === SHOPPING CART ===
// Product: Gaming Keyboard
// Price: $89.99
// Quantity: 3
// Subtotal: $269.97
// Discount (15%): $40.50
// Total: $229.47

func main() {
	productName := "Gaming Keyboard"
	price := "89.99"
	quantity := "3"
	discount := 16

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	discountFloat := float64(discount)
	totalPrice := priceFloat * float64(quantityInt)
	discountPrice := (totalPrice * discountFloat) / 100
	discountPercent := strconv.FormatFloat(discountFloat, 'f', 0, 64)
	totalPriceWithDiscount := totalPrice - discountPrice
	totalString := strconv.FormatFloat(totalPriceWithDiscount, 'f', 2, 64)
	result := fmt.Sprintf("Product: %s\n Price: %.2f\n Quantity: %d\n Subtotal: $%.2f\n Discount(%s%%): $%.2f\n Total: $%s", productName, priceFloat, quantityInt, totalPrice, discountPercent, discountPrice, totalString)
	fmt.Println(result)
}
