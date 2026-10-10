package main

import "fmt"

// Напиши программу для конвертации валют.
// Исходные данные:
// const usdToMdl = 17.25
// const eurToMdl = 20.10
// usd := 150.0
// eur := 80.0
// Что нужно сделать:
// Перевести USD в MDL.
// Перевести EUR в MDL.
// Посчитать общую сумму в MDL.
// Вывести результаты с точностью до двух знаков.
// Вывести тип каждой переменной через %T.
// Ожидаемый результат:
// === Currency Converter ===
// USD: 150.00
// EUR: 80.00
// USD to MDL: 2587.50
// EUR to MDL: 1608.00
// Total MDL: 4195.50
func main() {
	const usdToMdl = 17.25
	const eurToMdl = 20.10
	const divider = "---------------"

	usd := 100.0
	eur := 100.0

	usdToMdlAmount := usd * usdToMdl
	eurToMdlAmount := eur * eurToMdl
	totalMdl := usdToMdlAmount + eurToMdlAmount

	fmt.Println("=== Currency Converter ===")
	fmt.Println("USD Currency: ", usdToMdl)
	fmt.Println("EUR Currency: ", eurToMdl)
	fmt.Println(divider)
	fmt.Println("USD: ", usd)
	fmt.Println("EUR: ", eur)
	fmt.Println(divider)
	fmt.Printf("USD to MDL: %.2f\n", usdToMdlAmount)
	fmt.Printf("EUR to MDL: %.2f\n", eurToMdlAmount)
	fmt.Println("Total MDL:", totalMdl)
}
