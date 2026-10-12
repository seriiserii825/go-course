package main

import (
	"github.com/seriiserii825/go-course/05.slice/utils"
)

// Практика — Shopping Cart (Корзина покупок)
// Начальные данные:
// prices := []float64{19.99, 45.50, 12.00, 89.90}
// Задания
// Выведи все цены товаров через for range.
// Добавь два новых товара с ценами 25.50 и 15.00 через append().
// Удали второй товар из slice.
// Посчитай общую стоимость всех товаров.
// Найди самый дорогой товар.
// Выведи количество товаров через len().
// Выполнено 0 из 6
// Ожидаемый результат
// === SHOPPING CART ===
// Products:
// 1. $19.99
// 2. $12.00
// 3. $89.90
// 4. $25.50
// 5. $15.00
// Total products: 5
// Total price: $162.39
// Most expensive: $89.90
func main() {
	prices := []float64{19.99, 45.50, 12.00, 89.90}
	utils.PrintShopTitle()
	utils.PrintPrices(prices)
	utils.PrintProductsLen(prices)
	utils.PrintTotalPrice(prices)
	utils.PrintMostExpensive(prices)
}
