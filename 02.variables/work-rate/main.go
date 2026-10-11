package main

import "fmt"

// Задание 2. Калькулятор рабочего времени
// Уровень 3/5
// Представь, что пишешь CLI-приложение для учёта рабочего времени разработчика.
// const hourlyRate = 25.0
// hoursMonday := 8.5
// hoursTuesday := 7.0
// hoursWednesday := 9.5
// hoursThursday := 8.0
// hoursFriday := 6.5
// Нужно вычислить:
// - Общее количество отработанных часов.
// - Среднее количество часов в день.
// - Зарплату за неделю.
// - Зарплату за месяц, принимая 4 недели.
// - Сколько часов не хватает до стандартных 40 часов (или насколько план превышен).
// Вывод:
// === Developer Worklog ===
// Total hours: 39.50
// Average per day: 7.90
// Weekly salary: $987.50
// Monthly salary: $3950.00
// Difference from 40 hours: -0.50
// Дополнительная сложность: поменяй ставку с 25.0 на 30.0, изменив только одну константу. Вся программа должна автоматически пересчитать значения.
func main() {

	const divider = "========================"
	const days = 5
	const hourlyRate = 30.0
	const weekHours = 40
	hoursMonday := 8.5
	hoursTuesday := 7.0
	hoursWednesday := 9.5
	hoursThursday := 8.0
	hoursFriday := 6.5
	totalHours := hoursMonday + hoursTuesday + hoursWednesday + hoursThursday + hoursFriday
	averagePerDay := float64(totalHours) / float64(days)
	weeklySalary := totalHours * hourlyRate
	monthlySalary := weeklySalary * 4
	differenceWeekHours := totalHours - weekHours

	fmt.Println(divider)
	fmt.Printf("Hourly rate: %.2f\n", hourlyRate)
	fmt.Printf("Total hours: %.2f\n", totalHours)
	fmt.Printf("Average per day: %.2f\n", averagePerDay)
	fmt.Printf("Weekly Salary: %.2f\n", weeklySalary)
	fmt.Printf("Monthly Salary: %.2f\n", monthlySalary)
	fmt.Printf("Difference from %d hours: %.2f\n", weekHours, differenceWeekHours)
}
