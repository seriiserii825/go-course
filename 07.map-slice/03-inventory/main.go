package main

import "fmt"

// Смоделировать простой склад товаров через `map[string]int`, где ключ — название товара, значение — количество на складе.
// Реализовать функции:
// - приход товара (увеличить количество, если товара нет — добавить);
// - списание товара (уменьшить количество; если товара не хватает — вернуть ошибку `error`, количество не должно уходить в минус);
// - получение списка товаров, которых осталось меньше заданного порога (результат — `[]string`, отсортированный по алфавиту);
// - удаление товара со склада, если его количество равно нулю (использовать `delete`).
// ## Результат
// - Функция `Receive(stock map[string]int, name string, qty int)` — приход товара.
// - Функция `Ship(stock map[string]int, name string, qty int) error` — списание товара с проверкой остатка.
// - Функция `LowStock(stock map[string]int, threshold int) []string` — товары с остатком ниже порога, по алфавиту.
// - Функция `RemoveEmpty(stock map[string]int)` — удаляет из map товары с нулевым остатком.
// - `main()` с минимум 5 товарами, демонстрирующий приход, списание (включая попытку списать больше, чем есть — с обработкой ошибки), вывод товаров с низким остатком и очистку пустых позиций.
// create stock struct
type Stock struct {
	Items map[string]int
}

func main() {
	stock := map[string]int{
		"Apples":  50,
		"Bananas": 30,
		"Oranges": 20,
		"Grapes":  15,
		"Mangoes": 5,
	}
	printStock(stock)
	receive(stock, "Apples", 3)
	printStock(stock)
}
func printStock(stock map[string]int) {
	fmt.Println(stock)
}

func receive(stock map[string]int, name string, qty int) {
	stock[name] += qty
}
